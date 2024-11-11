// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"latihan/database"
	"latihan/repository"
	"latihan/service"
)

// Injectors from wire.go:

func InitializeService() *service.Service {
	db := database.ConnectDB()
	repositoryI := repository.NewRepository(db)
	serviceService := service.NewService(repositoryI)
	return serviceService
}
