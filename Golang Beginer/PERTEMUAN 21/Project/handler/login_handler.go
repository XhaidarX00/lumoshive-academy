package handler

import (
	"encoding/json"
	"fmt"
	"main/model"
	"main/utils"
	"net/http"
)

// curl -X POST http://localhost:8080/api/auth/login -H "Content-Type: application/json" -d '{"email":"admin1@example.com", "password":"password123"}'
func LoginHandler(w http.ResponseWriter, r *http.Request) {
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

	db, servicef := InitDBHandler(w)
	if db == nil {
		return
	}

	err = servicef.LoginService(&user)
	if err != nil {
		badResponse := model.Response{
			StatusCode: http.StatusBadRequest,
			Message:    "Account Not Found",
			Data:       nil,
		}
		json.NewEncoder(w).Encode(badResponse)
		return
	}

	token := servicef.Repo.GenerateTkn(user.ID, w)
	msg := fmt.Sprintf("Berhasil login, TOKEN ANDA : %s", token)
	utils.SuccesMessage(msg)

	response := model.Response{
		StatusCode: http.StatusOK,
		Message:    "Login success",
		Data:       user,
	}
	json.NewEncoder(w).Encode(response)
}
