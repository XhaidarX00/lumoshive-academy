package handler

// import (
// 	"be-golang-chapter-21/impleme-http-serve/model"
// 	"encoding/json"
// 	"net/http"
// )

// func RegisterHandler(w http.ResponseWriter, r *http.Request) {
// 	users := model.User{}
// 	err := json.NewDecoder(r.Body).Decode(&users)

// 	if err != nil {
// 		badResponse := model.Response{
// 			StatusCode: http.StatusBadRequest,
// 			Message:    "Error server",
// 			Data:       nil,
// 		}
// 		json.NewEncoder(w).Encode(badResponse)
// 		return
// 	}

// 	serviceUsers, db := InitHandler(w)
// 	if db == nil {
// 		return
// 	}
// 	defer db.Close()

// 	err = serviceUsers.LoginService(&users)
// 	if err != nil {
// 		badResponse := model.Response{
// 			StatusCode: http.StatusBadRequest,
// 			Message:    "Account Not Found",
// 			Data:       nil,
// 		}
// 		json.NewEncoder(w).Encode(badResponse)
// 		return
// 	}

// 	response := model.Response{
// 		StatusCode: http.StatusOK,
// 		Message:    "Login success",
// 		Data:       users,
// 	}
// 	json.NewEncoder(w).Encode(response)
// }
