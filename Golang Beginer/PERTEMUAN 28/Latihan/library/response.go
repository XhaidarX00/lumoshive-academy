package library

import (
	"encoding/json"
	"latihan/model/response"
	"net/http"
)

// OKRequest - Status 200 OK
func OKRequest(text string, data interface{}) response.Response {
	return response.Response{
		StatusCode: http.StatusOK,
		Message:    text,
		Data:       data,
	}
}

// CreatedRequest - Status 201 Created
func CreatedRequest(text string, data interface{}) response.Response {
	return response.Response{
		StatusCode: http.StatusCreated,
		Message:    text,
		Data:       data,
	}
}

// BadRequest - Status 400 Bad Request
func BadRequest(text string) response.ResponseError {
	return response.ResponseError{
		StatusCode: http.StatusBadRequest,
		Message:    text,
	}
}

// UnauthorizedRequest - Status 401 Unauthorized
func UnauthorizedRequest(text string) response.ResponseError {
	return response.ResponseError{
		StatusCode: http.StatusUnauthorized,
		Message:    text,
	}
}

// ForbiddenRequest - Status 403 Forbidden
func ForbiddenRequest(text string) response.ResponseError {
	return response.ResponseError{
		StatusCode: http.StatusForbidden,
		Message:    text,
	}
}

// NotFoundRequest - Status 404 Not Found
func NotFoundRequest(text string) response.ResponseError {
	return response.ResponseError{
		StatusCode: http.StatusNotFound,
		Message:    text,
	}
}

// ConflictRequest - Status 409 Conflict
func ConflictRequest(text string) response.ResponseError {
	return response.ResponseError{
		StatusCode: http.StatusConflict,
		Message:    text,
	}
}

// InternalServerError - Status 500 Internal Server Error
func InternalServerError(text string) response.ResponseError {
	return response.ResponseError{
		StatusCode: http.StatusInternalServerError,
		Message:    text,
	}
}

// ServiceUnavailableRequest - Status 503 Service Unavailable
func ServiceUnavailableRequest(text string) response.ResponseError {
	return response.ResponseError{
		StatusCode: http.StatusServiceUnavailable,
		Message:    text,
	}
}

func PageResponse(text string, limit, page, totalItems, totalPages int, data interface{}) response.PaginationResponse {
	return response.PaginationResponse{
		StatusCode: http.StatusOK,
		Message:    text,
		Page:       page,
		Limit:      limit,
		TotalItems: totalItems,
		TotalPages: totalPages,
		Data:       data,
	}
}

func JsonResponse(w http.ResponseWriter, response interface{}) {
	json.NewEncoder(w).Encode(response)
}

func ResponseToJson(w http.ResponseWriter, msg string, data interface{}) {
	if data != nil {
		response := OKRequest(msg, data)
		json.NewEncoder(w).Encode(response)

	} else {
		response := BadRequest(msg)
		json.NewEncoder(w).Encode(response)
	}
}
