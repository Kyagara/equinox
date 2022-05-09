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

	RateLimitedError = ErrorResponse{
		Status: Status{
			Message:    "Rate limited",
			StatusCode: http.StatusTooManyRequests,
		},
	}
)
