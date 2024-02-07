package errors

import "net/http"

const (
	API_ERROR_TYPE = "api_error"
)

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
		Type:   API_ERROR_TYPE,
		Title:  "Bad request",
		Detail: detail,
		Status: http.StatusBadRequest,
	}
}

func NotFoundError(detail string) ApiError {
	return ApiError{
		Type:   API_ERROR_TYPE,
		Title:  "Not found",
		Detail: detail,
		Status: http.StatusNotFound,
	}
}

func InternalServerError() ApiError {
	return ApiError{
		Type:   API_ERROR_TYPE,
		Title:  "Internal server error",
		Detail: "Internal server error",
		Status: http.StatusInternalServerError,
	}
}
