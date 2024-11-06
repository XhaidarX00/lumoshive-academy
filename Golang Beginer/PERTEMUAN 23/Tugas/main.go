package main

import (
	"main/database"
	"main/handler"
	"main/router"

	"main/service"
)

func main() {

	// Inisialisasi db dan service
	database.ConnectDB()
	service.NewService()

	// check token expired
	go service.ServiceF.CheckToken()

	// Inisialisasi template saat server dimulai
	handler.InitTemplates()

	// Inisialisasi Routes
	router.InitRoute()
}
