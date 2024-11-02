package model

import (
	_ "github.com/lib/pq"
)

type Product struct {
	ID       string  `json:"id"`
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Stock    int     `json:"stock"`
	Category string  `json:"category"`
	Location string  `json:"location"`
}

type PaginationResponse struct {
	StatusCode int       `json:"statusCode"`
	Message    string    `json:"message"`
	Page       int       `json:"page"`
	Limit      int       `json:"limit"`
	TotalItems int       `json:"total_items"`
	TotalPages int       `json:"total_pages"`
	Data       []Product `json:"data"`
}
