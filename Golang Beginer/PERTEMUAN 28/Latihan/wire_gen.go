// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"database/sql"
	"github.com/google/wire"
	"go.uber.org/zap"
	"latihan/controller"
	"latihan/database"
	"latihan/library"
	"latihan/repository"
	"latihan/service"
)

// Injectors from wire.go:

// 3 return object, func, error
func InitializeService() (Service, error) {
	db, err := database.ConnectDB()
	if err != nil {
		return Service{}, err
	}
	travel := repository.NewRepo(db)
	serviceService := service.NewService(travel)
	logger := library.InitLog()
	controllerTravel := controller.NewTravelHandelr(serviceService, logger)
	mainService := Service{
		Handler: controllerTravel,
		DB:      db,
		Logger:  logger,
	}
	return mainService, nil
}

// wire.go:

type Service struct {
	Handler *controller.Travel
	DB      *sql.DB
	Logger  *zap.Logger
}

var Servicset = wire.NewSet(database.ConnectDB, library.InitLog, repository.NewRepo, service.NewService, controller.NewTravelHandelr)