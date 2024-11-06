package main

import (
	"fmt"
	"main/database"
	"main/handler"
	"main/middleware"
	"main/service"
	"net/http"
)

func main() {
	database.ConnectDB()
	service.NewService()

	go service.ServiceF.CheckToken()

	serverMux := http.NewServeMux()

	// Setup authentication routes
	authMux := http.NewServeMux()
	// authMux.HandleFunc("POST /login", handler.LoginHandler)
	authMux.HandleFunc("POST /register", handler.RegisterHandler)
	authMux.HandleFunc("/", handler.RegistrationHandler)
	authMux.HandleFunc("/list-users", handler.ReadUsers)
	authMux.HandleFunc("/user-detail", handler.UserDetail)
	authMux.HandleFunc("/user-delete", handler.DeleteUsers)
	authMux.HandleFunc("/add-todo", handler.AddTodo)
	authMux.HandleFunc("/todo-list", handler.ReadTodos)
	authMux.HandleFunc("/change-status-todo", handler.ChangeTodoStatus)
	authMux.HandleFunc("/delete-todo", handler.DeleteTodo)

	// Setup resource routes with authentication and role middleware
	resourceMux := http.NewServeMux()
	// resourceMux.HandleFunc("GET /read", handler.ReadTaskHandler)
	resourceMux.HandleFunc("POST /create", handler.CreateTaskHandler)
	// resourceMux.HandleFunc("PUT /update", handler.UpdateTaskHandler)
	// resourceMux.HandleFunc("DELETE /delete", handler.DeleteTaskHandler)
	// resourceMux.HandleFunc("PUT /todos/{id}/status", handler.UpdateTaskStatusHandler)
	// resourceMux.HandleFunc("GET /users", handler.GetUsersHandler)
	// resourceMux.HandleFunc("GET /users/{id}", handler.GetUserDetailHandler)

	// Apply middleware chain
	protected := middleware.TokenMiddleware(
		middleware.RoleMiddleware(resourceMux),
	)

	// Inisialisasi template saat server dimulai
	handler.InitTemplates()

	// Mendaftarkan handler untuk file statis
	// fs := http.FileServer(http.Dir("./static"))
	// serverMux.Handle("/static", http.StripPrefix("/static/", fs))

	// API routes
	serverMux.Handle("/", authMux)
	serverMux.Handle("/api/", http.StripPrefix("/api/todos", protected))

	// Static file server
	// serverMux.Handle("/", middleware.CorsMiddleware(fs))

	fmt.Println("server started on port 8080")
	http.ListenAndServe(":8080", serverMux)
}
