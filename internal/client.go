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
	LogRequestFormat = "[%s '%s'] %s"
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

// Creates and sends a request using the parameters specified
func (c *InternalClient) SendRequest(method string, url string, endpoint string, v interface{}) error {
	// Creating a new *http.Request
	req, err := c.newRequest(method, fmt.Sprintf("%s%s", url, endpoint))

	if err != nil {
		return err
	}

	if c.debug {
		c.log.Info.Printf(LogRequestFormat, method, endpoint, "Requesting")
	}

	// Sending request
	res, err := c.http.Do(req)

	if err != nil {
		return err
	}

	defer res.Body.Close()

	// If the API returns with a 429, get the Retry-After header and retry after the specified time has passed
	if res.StatusCode == http.StatusTooManyRequests {
		retryAfter := res.Header.Get("Retry-After")

		seconds, err := strconv.Atoi(retryAfter)

		if err != nil {
			return err
		}

		if c.debug {
			c.log.Error.Printf(LogRequestFormat, method, endpoint, fmt.Sprintf("Too many requests, retrying in %ds", seconds))
		}

		time.Sleep(time.Duration(seconds) * time.Second)

		return c.SendRequest(method, url, endpoint, v)
	}

	// If the status code is lower than 200 or higher than 400
	if res.StatusCode < http.StatusOK || res.StatusCode >= http.StatusBadRequest {
		var errRes ErrorResponse

		if err = json.NewDecoder(res.Body).Decode(&errRes); err == nil {
			return fmt.Errorf("status code: %d, %v", errRes.Status.StatusCode, errRes.Status.Message)
		}

		return fmt.Errorf("unknown error, status code: %d", res.StatusCode)
	}

	// Decoding the json into the endpoint method response object
	if err = json.NewDecoder(res.Body).Decode(&v); err != nil {
		return err
	}

	if c.debug {
		c.log.Info.Printf(LogRequestFormat, method, endpoint, "Request successful")
	}

	return nil
}

// Creates a new *http.request and sets headers
func (c *InternalClient) newRequest(method string, url string) (*http.Request, error) {
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", "application/json; charset=utf-8")
	req.Header.Set("X-Riot-Token", c.key)

	return req, nil
}
