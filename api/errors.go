package api

import "net/http"

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
	NotFoundError = ErrorResponse{
		Status: Status{
			Message:    "Not Found",
			StatusCode: http.StatusNotFound,
		},
	}

	ForbiddenError = ErrorResponse{
		Status: Status{
			Message:    "Forbidden",
			StatusCode: http.StatusForbidden,
		},
	}
)
