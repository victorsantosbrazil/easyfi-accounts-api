package errors

import "net/http"

type ApiError struct {
	Type   string `json:"type"`
	Title  string `json:"title"`
	Detail string `json:"detail"`
	Status int    `json:"status"`
}

func (e ApiError) Error() string {
	return e.Detail
}

func BadRequestError(detail string) ApiError {
	return ApiError{
		Type:   "bad_request",
		Title:  "Bad Request",
		Detail: detail,
		Status: http.StatusBadRequest,
	}
}

func InternalServerError(detail string) ApiError {
	return ApiError{
		Type:   "internal_server",
		Title:  "Internal Server Error",
		Detail: detail,
		Status: http.StatusInternalServerError,
	}
}
