// package main

// import (
// 	"fmt"
// 	"html/template"
// 	"net/http"
// 	"path/filepath"
// )

// // renderTemplate merender template yang diminta berdasarkan parameter `tmpl` dengan `layout.html` sebagai template utama
// func renderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
// 	// Gabungkan `layout.html` dengan template yang spesifik
// 	templates, err := template.ParseFiles(
// 		filepath.Join("./templates", "layout.html"),
// 		filepath.Join("./templates", "navbar.html"),
// 		filepath.Join("./templates", tmpl),
// 	)
// 	if err != nil {
// 		http.Error(w, "Template tidak ditemukan atau gagal dirender", http.StatusInternalServerError)
// 		return
// 	}

// 	// Eksekusi template dengan `layout.html` sebagai template utama
// 	err = templates.ExecuteTemplate(w, "layout.html", data)
// 	if err != nil {
// 		http.Error(w, "Gagal menampilkan template", http.StatusInternalServerError)
// 		fmt.Println("Error executing template:", err) // Memperbaiki urutan error
// 	}
// }

// // Handler untuk halaman registrasi
// func registrationHandler(w http.ResponseWriter, r *http.Request) {
// 	renderTemplate(w, "registration.html", map[string]string{"Title": "Registration Page"})
// }

// // Handler untuk halaman todo
// func todoHandler(w http.ResponseWriter, r *http.Request) {
// 	renderTemplate(w, "todo-list.html", map[string]string{"Title": "Todo List"})
// }

// // Handler untuk halaman daftar pengguna
// func userHandler(w http.ResponseWriter, r *http.Request) {
// 	renderTemplate(w, "user.html", map[string]string{"Title": "User  List"})
// }

// func main() {
// 	http.HandleFunc("/register", registrationHandler)
// 	http.HandleFunc("/todos", todoHandler)
// 	http.HandleFunc("/users", userHandler)

// 	fs := http.FileServer(http.Dir("static"))
// 	http.Handle("/static/", http.StripPrefix("/static/", fs))

// 	// Menjalankan server di port 8080
// 	err := http.ListenAndServe(":8080", nil)
// 	if err != nil {
// 		panic("Server gagal berjalan: " + err.Error())
// 	}
// }

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
	fs := http.FileServer(http.Dir("./static"))
	// staticHandler := http.StripPrefix("/", fs)
	serverMux.Handle("/", fs)

	// Register all routes to main mux
	// API routes
	serverMux.Handle("/api/auth/", http.StripPrefix("/api/auth", authMux))
	serverMux.Handle("/api/", http.StripPrefix("/api/todos", protected))

	// Static file server
	// serverMux.Handle("/", middleware.CorsMiddleware(fs))

	fmt.Println("server started on port 8080")
	http.ListenAndServe(":8080", serverMux)

}
