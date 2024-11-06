package main

import (
	"log"
	"main/database"
	"main/handler"
	"main/repository"
	"main/router"
	"main/service"
	"main/utils"
	"net/http"
)

func main() {
	db, err := database.InitDB()
	if err != nil {
		utils.ErrorMessage(err.Error())
		utils.ExitProgram()
	}

	defer db.Close()

	userRepo := repository.NewUserRepository(db)
	if userRepo == nil {
		log.Fatal("Failed to create user repository")
	}

	// Initialize service
	authService := service.NewAuthService(userRepo)
	if authService == nil {
		log.Fatal("Failed to create auth service")
	}

	// Initialize handler
	authHandler := handler.NewAuthHandler(authService)
	if authHandler == nil {
		log.Fatal("Failed to create auth handler")
	}

	// Setup router
	mux := router.SetupRouter(authHandler)
	if mux == nil {
		log.Fatal("Failed to setup router")
	}

	// Start server
	log.Println("Server starting on :8000")
	server := &http.Server{
		Addr:    ":8000",
		Handler: mux,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Fatal("Server failed to start:", err)
	}
}
