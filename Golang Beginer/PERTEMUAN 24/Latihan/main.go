package main

import (
	handler "latihan/controller"
	"latihan/database"
	"latihan/router"
	"latihan/service"
)

func main() {
	// init db
	database.ConnectDB()

	// init service
	service.NewService()

	// init template
	handler.InitTemplates()

	// init routes
	router.InitRoute()

}
