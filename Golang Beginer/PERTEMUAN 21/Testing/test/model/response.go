package model

type RegistrationResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

type LoginResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Token   string `json:"token,omitempty"`
}

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
