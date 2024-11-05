package model

type Response struct {
	StatusCode int
	Message    string
	Data       interface{}
}

type ResponseError struct {
	StatusCode int
	Message    string
}

type PaginationResponse struct {
	StatusCode int
	Message    string
	Page       int
	Limit      int
	TotalItems int
	TotalPages int
	Data       interface{}
}
