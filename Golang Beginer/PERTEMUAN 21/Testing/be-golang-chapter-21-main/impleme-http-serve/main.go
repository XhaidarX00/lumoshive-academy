package main

import (
	"be-golang-chapter-21/impleme-http-serve/handler"
	"be-golang-chapter-21/impleme-http-serve/middleware"
	"fmt"
	"net/http"
)

func main() {
	serverMux := http.NewServeMux()

	authMux := http.NewServeMux()
	authMux.HandleFunc("POST /login", handler.LoginHandler)

	resourceMux := http.NewServeMux()
	resourceMux.HandleFunc("GET /customer_detail", handler.GetCustomerByID)

	role := middleware.Role(resourceMux)
	middleware := middleware.Middleware(role)

	serverMux.Handle("/", authMux)
	serverMux.Handle("/customer/", http.StripPrefix("/customer", middleware))

	fmt.Println("server started on port 8080")
	http.ListenAndServe(":8080", serverMux)
}
