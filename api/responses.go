package api

import (
	"net/http"
)

type HTTPErrorResponse struct {
	Status Status `json:"status"`
}

type Status struct {
	Message    string `json:"message"`
	StatusCode int    `json:"status_code"`
}

func (e HTTPErrorResponse) Error() string {
	return e.Status.Message
}

var (
	ErrBadRequest = HTTPErrorResponse{
		Status: Status{
			Message:    "Bad Request",
			StatusCode: http.StatusBadRequest,
		},
	}
	ErrUnauthorized = HTTPErrorResponse{
		Status: Status{
			Message:    "Unauthorized",
			StatusCode: http.StatusUnauthorized,
		},
	}
	ErrForbidden = HTTPErrorResponse{
		Status: Status{
			Message:    "Forbidden",
			StatusCode: http.StatusForbidden,
		},
	}
	ErrNotFound = HTTPErrorResponse{
		Status: Status{
			Message:    "Not Found",
			StatusCode: http.StatusNotFound,
		},
	}
	ErrMethodNotAllowed = HTTPErrorResponse{
		Status: Status{
			Message:    "Method not allowed",
			StatusCode: http.StatusMethodNotAllowed,
		},
	}
	ErrUnsupportedMediaType = HTTPErrorResponse{
		Status: Status{
			Message:    "Unsupported media type",
			StatusCode: http.StatusUnsupportedMediaType,
		},
	}
	ErrTooManyRequests = HTTPErrorResponse{
		Status: Status{
			Message:    "Rate limited",
			StatusCode: http.StatusTooManyRequests,
		},
	}
	ErrInternalServer = HTTPErrorResponse{
		Status: Status{
			Message:    "Internal server error",
			StatusCode: http.StatusInternalServerError,
		},
	}
	ErrBadGateway = HTTPErrorResponse{
		Status: Status{
			Message:    "Bad gateway",
			StatusCode: http.StatusBadGateway,
		},
	}
	ErrServiceUnavailable = HTTPErrorResponse{
		Status: Status{
			Message:    "Service unavailable",
			StatusCode: http.StatusServiceUnavailable,
		},
	}
	ErrGatewayTimeout = HTTPErrorResponse{
		Status: Status{
			Message:    "Gateway timeout",
			StatusCode: http.StatusGatewayTimeout,
		},
	}
	StatusCodeToError = map[int]HTTPErrorResponse{
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
