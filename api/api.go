// Package used to share common constants and structs.
package api

import (
	"errors"
	"net/http"

	"github.com/rs/zerolog"
)

// Base API URLs formats.
const (
	RIOT_API_BASE_URL_FORMAT = ".api.riotgames.com"
)

// EquinoxRequest represents a request to the Riot API, its a struct that contains all information about a request.
type EquinoxRequest struct {
	Logger   zerolog.Logger
	Route    string
	Request  *http.Request
	URL      string
	MethodID string
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

func StatusCodeToError(statusCode int) error {
	switch statusCode {
	case http.StatusBadRequest:
		return ErrBadRequest
	case http.StatusUnauthorized:
		return ErrUnauthorized
	case http.StatusForbidden:
		return ErrForbidden
	case http.StatusNotFound:
		return ErrNotFound
	case http.StatusMethodNotAllowed:
		return ErrMethodNotAllowed
	case http.StatusUnsupportedMediaType:
		return ErrUnsupportedMediaType
	case http.StatusTooManyRequests:
		return ErrTooManyRequests
	case http.StatusInternalServerError:
		return ErrInternalServer
	case http.StatusBadGateway:
		return ErrBadGateway
	case http.StatusServiceUnavailable:
		return ErrServiceUnavailable
	case http.StatusGatewayTimeout:
		return ErrGatewayTimeout
	default:
		return nil
	}
}
