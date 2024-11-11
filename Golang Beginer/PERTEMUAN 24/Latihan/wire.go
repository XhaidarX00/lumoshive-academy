//go:build wireinject
// +build wireinject

package main

import (
	"latihan/database"
	"latihan/repository"
	"latihan/service"

	"github.com/google/wire"
)

func InitializeService() *service.Service {
	wire.Build(database.ConnectDB, service.NewService, repository.NewRepository)
	return nil
}
