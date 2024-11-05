package main

import (
	"be-golang-chapter-21/impleme-http-serve/handler"
	"be-golang-chapter-21/impleme-http-serve/middleware"
	"fmt"
	"net/http"
)

// TODO LIST https://drive.google.com/file/d/1ApiUhpm0QWrNneCsJXCEjaDTjT1c_A4K/view?usp=sharing

func main() {
	serverMux := http.NewServeMux() // general masuk listen

	authMux := http.NewServeMux() // handler login not implemen midle
	authMux.HandleFunc("POST /login", handler.LoginHandler)

	resourceMux := http.NewServeMux() //
	resourceMux.HandleFunc("GET /customer_detail", handler.GetUsersByID)

	role := middleware.Role(resourceMux)
	middleware := middleware.TokenMiddleware(role) // auth token

	serverMux.Handle("/", authMux)
	serverMux.Handle("/customer/", http.StripPrefix("/customer", middleware))

	fmt.Println("server started on port 8080")
	http.ListenAndServe(":8080", serverMux)
}
