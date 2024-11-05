package handler

import (
	"be-golang-chapter-21/impleme-http-serve/database"
	"be-golang-chapter-21/impleme-http-serve/model"
	"be-golang-chapter-21/impleme-http-serve/repository"
	"be-golang-chapter-21/impleme-http-serve/service"
	"database/sql"
	"encoding/json"
	"net/http"
)

func InitHandler(w http.ResponseWriter) (service.UserService, *sql.DB) {
	db, err := database.InitDB()
	if err != nil {
		badResponse := model.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		}
		json.NewEncoder(w).Encode(badResponse)
		return service.UserService{}, nil
	}

	repo := repository.NewUserRepository(db)
	serviceCustomer := service.NewUserService(repo)

	return serviceCustomer, db
}
