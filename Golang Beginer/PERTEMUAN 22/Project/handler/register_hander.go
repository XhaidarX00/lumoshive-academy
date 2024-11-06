package handler

import (
	"encoding/json"
	"fmt"
	"main/model"
	"main/utils"
	"net/http"
)

// curl -X POST http://localhost:8080/api/auth/register -H "Content-Type: application/json" -d '{"username":"admin3", "password":"password789", "email":"admin3@example.com", "role":"admin"}'
func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	user := model.User{}
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		badResponse := model.Response{
			StatusCode: http.StatusBadRequest,
			Message:    "Error server",
			Data:       nil,
		}
		json.NewEncoder(w).Encode(badResponse)
		return
	}

	fmt.Printf("Regis Data : %v\n", user)

	if len(user.Password) < 8 {
		badResponse := model.ResponseError{
			StatusCode: http.StatusBadRequest,
			Message:    "must be at lease 8 character in password",
		}
		json.NewEncoder(w).Encode(badResponse)
		return
	}

	db, servicef := InitDBHandler(w)
	if db == nil {
		return
	}

	err = servicef.RegisterService(&user)
	if err != nil {
		badResponse := model.ResponseError{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
		}
		json.NewEncoder(w).Encode(badResponse)
		return
	}

	msg := fmt.Sprintf("Berhasil Register, Data : %v\n", user)
	utils.SuccesMessage(msg)

	response := model.Response{
		StatusCode: http.StatusCreated,
		Message:    "Register success",
		Data:       nil,
	}
	json.NewEncoder(w).Encode(response)
}
