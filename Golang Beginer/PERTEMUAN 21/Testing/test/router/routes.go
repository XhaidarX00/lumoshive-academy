package router

import (
	"main/handler"
	"net/http"
)

func SetupRouter(authHandler *handler.AuthHandler) *http.ServeMux {
	mux := http.NewServeMux()

	// Serve static files
	fs := http.FileServer(http.Dir("./view"))
	mux.Handle("/", fs)

	// API endpoints
	mux.HandleFunc("/api/login", authHandler.Login)
	mux.HandleFunc("/getid", authHandler.GetUserById)

	return mux
}
