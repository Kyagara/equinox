package api

import "net/http"

type PlainTextResponse struct {
	Response interface{} `json:"response"`
}

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
	BadRequestError = ErrorResponse{
		Status: Status{
			Message:    "Bad Request",
			StatusCode: http.StatusBadRequest,
		},
	}

	UnauthorizedError = ErrorResponse{
		Status: Status{
			Message:    "Unauthorized",
			StatusCode: http.StatusUnauthorized,
		},
	}

	ForbiddenError = ErrorResponse{
		Status: Status{
			Message:    "Forbidden",
			StatusCode: http.StatusForbidden,
		},
	}

	NotFoundError = ErrorResponse{
		Status: Status{
			Message:    "Not Found",
			StatusCode: http.StatusNotFound,
		},
	}

	MethodNotAllowedError = ErrorResponse{
		Status: Status{
			Message:    "Method not allowed",
			StatusCode: http.StatusMethodNotAllowed,
		},
	}

	UnsupportedMediaTypeError = ErrorResponse{
		Status: Status{
			Message:    "Unsupported media type",
			StatusCode: http.StatusUnsupportedMediaType,
		},
	}

	RateLimitedError = ErrorResponse{
		Status: Status{
			Message:    "Rate limited",
			StatusCode: http.StatusTooManyRequests,
		},
	}

	InternalServerError = ErrorResponse{
		Status: Status{
			Message:    "Internal server error",
			StatusCode: http.StatusInternalServerError,
		},
	}

	BadGatewayError = ErrorResponse{
		Status: Status{
			Message:    "Bad gateway",
			StatusCode: http.StatusBadGateway,
		},
	}

	ServiceUnavailableError = ErrorResponse{
		Status: Status{
			Message:    "Service unavailable",
			StatusCode: http.StatusServiceUnavailable,
		},
	}

	GatewayTimeoutError = ErrorResponse{
		Status: Status{
			Message:    "Gateway timeout",
			StatusCode: http.StatusGatewayTimeout,
		},
	}

	StatusCodeToError = map[int]ErrorResponse{
		http.StatusBadRequest:           BadRequestError,
		http.StatusUnauthorized:         UnauthorizedError,
		http.StatusForbidden:            ForbiddenError,
		http.StatusNotFound:             NotFoundError,
		http.StatusMethodNotAllowed:     MethodNotAllowedError,
		http.StatusUnsupportedMediaType: UnsupportedMediaTypeError,
		http.StatusTooManyRequests:      RateLimitedError,
		http.StatusInternalServerError:  InternalServerError,
		http.StatusBadGateway:           BadGatewayError,
		http.StatusServiceUnavailable:   ServiceUnavailableError,
		http.StatusGatewayTimeout:       GatewayTimeoutError,
	}
)
