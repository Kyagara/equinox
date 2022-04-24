package internal

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

type InternalClient struct {
	key   string
	debug bool
	http  *http.Client
	log   *Logger
}

const (
	LogRequestFormat = "[Method: '%s' | Query: '%v'] %s"
)

// Returns a new client using the API key provided
func NewInternalClient(key string, debug bool) *InternalClient {
	return &InternalClient{
		key:   key,
		debug: debug,
		http:  &http.Client{Timeout: 5 * time.Second},
		log:   NewLogger(),
	}
}

type ErrorResponse struct {
	Status struct {
		Message    string `json:"message"`
		StatusCode int    `json:"status_code"`
	} `json:"status"`
}

func (c *InternalClient) SendRequest(req *http.Request, method string, v interface{}) error {
	req.Header.Set("Accept", "application/json; charset=utf-8")
	req.Header.Set("X-Riot-Token", c.key)

	if c.debug {
		c.log.Info.Printf(LogRequestFormat, method, req.URL.Query(), "Requesting")
	}

	res, err := c.http.Do(req)

	if err != nil {
		return err
	}

	defer res.Body.Close()

	if res.StatusCode == http.StatusTooManyRequests {
		retryAfter := res.Header.Get("Retry-After")

		seconds, err := strconv.Atoi(retryAfter)

		if err != nil {
			return err
		}

		if c.debug {
			c.log.Error.Printf(LogRequestFormat, method, req.URL.Query(), fmt.Sprintf("Too many requests, retrying in %ds", seconds))
		}

		time.Sleep(time.Duration(seconds) * time.Second)

		return c.SendRequest(req, method, v)
	}

	if res.StatusCode < http.StatusOK || res.StatusCode >= http.StatusBadRequest {
		var errRes ErrorResponse

		if err = json.NewDecoder(res.Body).Decode(&errRes); err == nil {
			return fmt.Errorf("status code: %d, %v", errRes.Status.StatusCode, errRes.Status.Message)
		}

		return fmt.Errorf("unknown error, status code: %d", res.StatusCode)
	}

	if err = json.NewDecoder(res.Body).Decode(&v); err != nil {
		return err
	}

	if c.debug {
		c.log.Info.Printf(LogRequestFormat, method, req.URL.Query(), "Request successful")
	}

	return nil
}
