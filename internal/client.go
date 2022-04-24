package internal

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type InternalClient struct {
	apiKey string
	http   *http.Client
}

// Returns a new client using the API key provided
func NewClient(key string) *InternalClient {
	return &InternalClient{
		apiKey: key,
		http:   &http.Client{Timeout: time.Minute},
	}
}

type ErrorResponse struct {
	Status struct {
		Message    string `json:"message"`
		StatusCode int    `json:"status_code"`
	} `json:"status"`
}

func (c *InternalClient) SendRequest(req *http.Request, v interface{}) error {
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("Accept", "application/json; charset=utf-8")
	req.Header.Set("X-Riot-Token", c.apiKey)

	res, err := c.http.Do(req)

	if err != nil {
		return err
	}

	defer res.Body.Close()

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

	return nil
}
