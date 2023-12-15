// Package used to share common constants and structs.
package api

import (
	"errors"
	"net/http"

	"github.com/rs/zerolog"
)

// Base API URLs formats.
const (
	RIOT_API_BASE_URL_FORMAT = "https://%s.api.riotgames.com%s"
	D_DRAGON_BASE_URL_FORMAT = "https://ddragon.leagueoflegends.com%s%s"
	C_DRAGON_BASE_URL_FORMAT = "https://cdn.communitydragon.org%s%s"
)

// EquinoxRequest represents a request to the Riot API and CDNs, its a struct that contains all information about a request.
type EquinoxRequest struct {
	Logger   zerolog.Logger
	Route    any
	Request  *http.Request
	URL      string
	MethodID string
	IsCDN    bool
}

var (
	ErrBadRequest           = errors.New("bad request")
	ErrUnauthorized         = errors.New("unauthorized")
	ErrForbidden            = errors.New("forbidden")
	ErrNotFound             = errors.New("not found")
	ErrMethodNotAllowed     = errors.New("method not allowed")
	ErrUnsupportedMediaType = errors.New("unsupported media type")
	ErrTooManyRequests      = errors.New("too many requests")
	ErrInternalServer       = errors.New("internal server error")
	ErrBadGateway           = errors.New("bad gateway")
	ErrServiceUnavailable   = errors.New("service unavailable")
	ErrGatewayTimeout       = errors.New("gateway timeout")
)

var (
	StatusCodeToError = map[int]error{
		http.StatusBadRequest:           ErrBadRequest,
		http.StatusUnauthorized:         ErrUnauthorized,
		http.StatusForbidden:            ErrForbidden,
		http.StatusNotFound:             ErrNotFound,
		http.StatusMethodNotAllowed:     ErrMethodNotAllowed,
		http.StatusUnsupportedMediaType: ErrUnsupportedMediaType,
		http.StatusTooManyRequests:      ErrTooManyRequests,
		http.StatusInternalServerError:  ErrInternalServer,
		http.StatusBadGateway:           ErrBadGateway,
		http.StatusServiceUnavailable:   ErrServiceUnavailable,
		http.StatusGatewayTimeout:       ErrGatewayTimeout,
	}
)
