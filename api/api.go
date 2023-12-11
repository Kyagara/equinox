// Package used to share common constants and structs.
package api

import (
	"net/http"

	"github.com/rs/zerolog"
)

// EquinoxRequest represents a request to the Riot API and CDNs, its a struct that contains all information about a request.
type EquinoxRequest struct {
	Logger   zerolog.Logger
	Route    any
	Request  *http.Request
	MethodID string
	Retries  int
	IsCDN    bool
}
