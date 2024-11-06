package handler

import (
	"be-golang-chapter-21/impleme-http-serve/database"
	"be-golang-chapter-21/impleme-http-serve/model"
	"be-golang-chapter-21/impleme-http-serve/repository"
	"be-golang-chapter-21/impleme-http-serve/service"
	"encoding/json"
	"net/http"
	"strconv"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	customer := model.Customer{}
	err := json.NewDecoder(r.Body).Decode(&customer)

	if err != nil {
		badResponse := model.Response{
			StatusCode: http.StatusBadRequest,
			Message:    "Error server",
			Data:       nil,
		}
		json.NewEncoder(w).Encode(badResponse)
		return
	}

	db, err := database.InitDB()
	if err != nil {
		badResponse := model.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		}
		json.NewEncoder(w).Encode(badResponse)
		return
	}
	defer db.Close()

	repo := repository.NewCustomerRepository(db)
	serviceCustomer := service.NewCustomerService(repo)

	err = serviceCustomer.LoginService(customer)
	if err != nil {
		badResponse := model.Response{
			StatusCode: http.StatusBadRequest,
			Message:    "Account Not Found",
			Data:       nil,
		}
		json.NewEncoder(w).Encode(badResponse)
		return
	}

	response := model.Response{
		StatusCode: http.StatusOK,
		Message:    "Login success",
		Data:       customer,
	}
	json.NewEncoder(w).Encode(response)
}

func GetCustomerByID(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	id := query.Get("id")

	id_int, _ := strconv.Atoi(id)

	db, err := database.InitDB()
	if err != nil {
		badResponse := model.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		}
		json.NewEncoder(w).Encode(badResponse)
		return
	}
	defer db.Close()

	repo := repository.NewCustomerRepository(db)
	serviceCustomer := service.NewCustomerService(repo)

	customer, err := serviceCustomer.CustomerByID(id_int)
	if err != nil {
		badResponse := model.Response{
			StatusCode: http.StatusBadRequest,
			Message:    "Account Not Found",
			Data:       nil,
		}
		json.NewEncoder(w).Encode(badResponse)
		return
	}

	response := model.Response{
		StatusCode: http.StatusOK,
		Message:    "success",
		Data:       customer,
	}
	json.NewEncoder(w).Encode(response)
}
