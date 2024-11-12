//go:build wireinject
// +build wireinject

package main

import (
	"database/sql"
	"latihan/controller"
	"latihan/database"
	"latihan/library"
	middlewaree "latihan/middleware"
	"latihan/repository"
	"latihan/service"

	"github.com/google/wire"
)

func InitializeService() (*controller.Controller, *sql.DB, error) {
	wire.Build(
		database.ConnectDB,
		library.InitLog,
		middlewaree.NewMiddleware,
		service.NewService,
		repository.NewRepository,
		controller.NewController,
	)

	return &controller.Controller{}, &sql.DB{}, error
}
