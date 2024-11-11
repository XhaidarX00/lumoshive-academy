package main

import (
	payments "latihan/controller/payment"
	"latihan/router"
)

func main() {
	// init db
	// database.ConnectDB()

	// init service
	// service.NewService()

	// init template
	// controller.InitTemplates()

	payments.Svc = InitializeService()

	// init routes
	router.InitRoute()
}
