package equinox

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Equinox struct {
	apikey string
	http   *http.Client
	LOL    *LOLClient
}

// Returns a new Equinox client using the API key provided
func NewClient(key string) *Equinox {
	client := &Equinox{
		apikey: key,
		http:   &http.Client{Timeout: time.Minute},
	}

	// Is there a better way to pass the clients fields without doing this?
	// I am passing the client to the LOLClient struct, this way I can call the sendRequest method inside endpoint methods
	client.LOL = NewLOLClient(client)

	return client
}

type ErrorResponse struct {
	Status struct {
		Message    string `json:"message"`
		StatusCode int    `json:"status_code"`
	} `json:"status"`
}

func (c *Equinox) sendRequest(req *http.Request, v interface{}) error {
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("Accept", "application/json; charset=utf-8")
	req.Header.Set("X-Riot-Token", c.apikey)

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
