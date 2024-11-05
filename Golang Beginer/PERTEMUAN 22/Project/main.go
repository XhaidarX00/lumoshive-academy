package main

import (
	"fmt"
	"main/handler"
	"main/middleware"
	"main/service"
	"net/http"
)

func main() {
	go service.CheckToken()
	serverMux := http.NewServeMux()

	// Setup authentication routes
	authMux := http.NewServeMux()
	authMux.HandleFunc("POST /login", handler.LoginHandler)
	authMux.HandleFunc("POST /register", handler.RegisterHandler)

	// Setup resource routes with authentication and role middleware
	resourceMux := http.NewServeMux()
	resourceMux.HandleFunc("GET /read", handler.ReadTaskHandler)
	resourceMux.HandleFunc("POST /create", handler.CreateTaskHandler)
	resourceMux.HandleFunc("PUT /update", handler.UpdateTaskHandler)
	resourceMux.HandleFunc("DELETE /delete", handler.DeleteTaskHandler)
	// resourceMux.HandleFunc("PUT /todos/{id}/status", handler.UpdateTaskStatusHandler)
	// resourceMux.HandleFunc("GET /users", handler.GetUsersHandler)
	// resourceMux.HandleFunc("GET /users/{id}", handler.GetUserDetailHandler)

	// Apply middleware chain
	protected := middleware.TokenMiddleware(
		middleware.RoleMiddleware(
			middleware.CorsMiddleware(resourceMux),
		),
	)

	// Setup file server for static files
	// fs := http.FileServer(http.Dir("./view"))
	// staticHandler := http.StripPrefix("/", fs)

	// Register all routes to main mux
	// API routes
	serverMux.Handle("/api/auth/", http.StripPrefix("/api/auth", authMux))
	serverMux.Handle("/api/", http.StripPrefix("/api/todos", protected))

	// Static file server
	// serverMux.Handle("/", middleware.CorsMiddleware(staticHandler))

	fmt.Println("server started on port 8080")
	http.ListenAndServe(":8080", serverMux)

}
