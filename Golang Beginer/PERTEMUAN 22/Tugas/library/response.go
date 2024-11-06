package library

import (
	"encoding/json"
	"main/model"
	"net/http"
)

// OKRequest - Status 200 OK
func OKRequest(response string, data interface{}) model.Response {
	return model.Response{
		StatusCode: http.StatusOK,
		Message:    response,
		Data:       data,
	}
}

// CreatedRequest - Status 201 Created
func CreatedRequest(response string, data interface{}) model.Response {
	return model.Response{
		StatusCode: http.StatusCreated,
		Message:    response,
		Data:       data,
	}
}

// BadRequest - Status 400 Bad Request
func BadRequest(response string) model.ResponseError {
	return model.ResponseError{
		StatusCode: http.StatusBadRequest,
		Message:    response,
	}
}

// UnauthorizedRequest - Status 401 Unauthorized
func UnauthorizedRequest(response string) model.ResponseError {
	return model.ResponseError{
		StatusCode: http.StatusUnauthorized,
		Message:    response,
	}
}

// ForbiddenRequest - Status 403 Forbidden
func ForbiddenRequest(response string) model.ResponseError {
	return model.ResponseError{
		StatusCode: http.StatusForbidden,
		Message:    response,
	}
}

// NotFoundRequest - Status 404 Not Found
func NotFoundRequest(response string) model.ResponseError {
	return model.ResponseError{
		StatusCode: http.StatusNotFound,
		Message:    response,
	}
}

// ConflictRequest - Status 409 Conflict
func ConflictRequest(response string) model.ResponseError {
	return model.ResponseError{
		StatusCode: http.StatusConflict,
		Message:    response,
	}
}

// InternalServerError - Status 500 Internal Server Error
func InternalServerError(response string) model.ResponseError {
	return model.ResponseError{
		StatusCode: http.StatusInternalServerError,
		Message:    response,
	}
}

// ServiceUnavailableRequest - Status 503 Service Unavailable
func ServiceUnavailableRequest(response string) model.ResponseError {
	return model.ResponseError{
		StatusCode: http.StatusServiceUnavailable,
		Message:    response,
	}
}

func PageResponse(response string, limit, page, totalItems, totalPages int, data interface{}) model.PaginationResponse {
	return model.PaginationResponse{
		StatusCode: http.StatusOK,
		Message:    response,
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
