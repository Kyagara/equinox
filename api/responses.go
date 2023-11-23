package api

import (
	"errors"
	"net/http"
)

type ErrorResponse struct {
	Status Status `json:"status"`
}

type Status struct {
	Message    string `json:"message"`
	StatusCode int    `json:"status_code"`
}

func (e ErrorResponse) Error() string {
	return e.Status.Message
}

var (
	ErrRetryAfterHeaderNotFound = errors.New("rate limited but no Retry-After header was found, stopping")
	ErrBadRequest               = ErrorResponse{
		Status: Status{
			Message:    "Bad Request",
			StatusCode: http.StatusBadRequest,
		},
	}
	ErrUnauthorized = ErrorResponse{
		Status: Status{
			Message:    "Unauthorized",
			StatusCode: http.StatusUnauthorized,
		},
	}
	ErrForbidden = ErrorResponse{
		Status: Status{
			Message:    "Forbidden",
			StatusCode: http.StatusForbidden,
		},
	}
	ErrNotFound = ErrorResponse{
		Status: Status{
			Message:    "Not Found",
			StatusCode: http.StatusNotFound,
		},
	}
	ErrMethodNotAllowed = ErrorResponse{
		Status: Status{
			Message:    "Method not allowed",
			StatusCode: http.StatusMethodNotAllowed,
		},
	}
	ErrUnsupportedMediaType = ErrorResponse{
		Status: Status{
			Message:    "Unsupported media type",
			StatusCode: http.StatusUnsupportedMediaType,
		},
	}
	ErrTooManyRequests = ErrorResponse{
		Status: Status{
			Message:    "Rate limited",
			StatusCode: http.StatusTooManyRequests,
		},
	}
	ErrInternalServer = ErrorResponse{
		Status: Status{
			Message:    "Internal server error",
			StatusCode: http.StatusInternalServerError,
		},
	}
	ErrBadGateway = ErrorResponse{
		Status: Status{
			Message:    "Bad gateway",
			StatusCode: http.StatusBadGateway,
		},
	}
	ErrServiceUnavailable = ErrorResponse{
		Status: Status{
			Message:    "Service unavailable",
			StatusCode: http.StatusServiceUnavailable,
		},
	}
	ErrGatewayTimeout = ErrorResponse{
		Status: Status{
			Message:    "Gateway timeout",
			StatusCode: http.StatusGatewayTimeout,
		},
	}
	StatusCodeToError = map[int]ErrorResponse{
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
