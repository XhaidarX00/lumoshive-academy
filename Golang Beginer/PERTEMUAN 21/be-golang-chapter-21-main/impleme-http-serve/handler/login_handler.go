package handler

import (
	"be-golang-chapter-21/impleme-http-serve/model"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	users := model.User{}
	err := json.NewDecoder(r.Body).Decode(&users)
	fmt.Printf("%v\n", users)
	if err != nil {
		badResponse := model.Response{
			StatusCode: http.StatusBadRequest,
			Message:    "Error server",
			Data:       nil,
		}
		json.NewEncoder(w).Encode(badResponse)
		return
	}

	serviceUsers, db := InitHandler(w)
	if db == nil {
		return
	}
	defer db.Close()

	err = serviceUsers.LoginService(&users)
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
		Data:       users,
	}
	json.NewEncoder(w).Encode(response)
}

func GetUsersByID(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	id := query.Get("id")

	id_int, _ := strconv.Atoi(id)
	serviceUsers, db := InitHandler(w)
	if db == nil {
		return
	}
	defer db.Close()

	// Users, err := serviceUsers.UserByID(id)
	Users, err := serviceUsers.UserByID(id_int)
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
		Data:       Users,
	}
	json.NewEncoder(w).Encode(response)
}
