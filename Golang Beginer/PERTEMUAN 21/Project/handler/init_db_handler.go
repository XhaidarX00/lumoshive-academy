package handler

import (
	"database/sql"
	"main/database"
	"main/repository"
	"main/service"
	"net/http"
)

func InitDBHandler(w http.ResponseWriter) (*sql.DB, service.Service) {
	db, err := database.InitDB(w)
	if err != nil {
		return nil, service.Service{}
	}

	repof := repository.NewRepository(db)
	servicef := service.NewService(repof)

	return db, servicef
}
