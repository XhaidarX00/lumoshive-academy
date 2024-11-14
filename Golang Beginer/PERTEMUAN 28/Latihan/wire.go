//go:build wireinject
// +build wireinject

package main

import (
	"database/sql"
	"latihan/controller"
	"latihan/database"
	"latihan/library"
	"latihan/repository"
	"latihan/service"

	"github.com/google/wire"
	"go.uber.org/zap"
)

type Service struct {
	Handler *controller.Travel
	DB      *sql.DB
	Logger  *zap.Logger
}

var Servicset = wire.NewSet(
	database.ConnectDB,
	library.InitLog,
	repository.NewRepo,
	service.NewService,
	controller.NewTravelHandelr,
)

// 3 return object, func, error
func InitializeService() (Service, error) {
	wire.Build(
		Servicset,
		wire.Struct(new(Service), "*"),
	)

	return Service{}, nil
}
