package response

type RegistrationResponse struct {
	Success bool
	Message string
}

type LoginResponse struct {
	Success bool
	Message string
	Token   string
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
