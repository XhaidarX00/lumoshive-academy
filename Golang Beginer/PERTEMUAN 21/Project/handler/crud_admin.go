package handler

import (
	"encoding/json"
	"fmt"
	"main/model"
	"main/utils"
	"net/http"
)

// curl -X POST http://localhost:8080/api/todos/create -H "token:token_admin_1" -H "Content-Type: application/json" -d '{"user_id":1, "title":"project golang", "description":"membuat api endpoint todo list"}'
func CreateTaskHandler(w http.ResponseWriter, r *http.Request) {
	task := model.Task{}
	err := json.NewDecoder(r.Body).Decode(&task)

	if err != nil {
		badResponse := model.ResponseError{
			StatusCode: http.StatusBadRequest,
			Message:    "Error server",
		}
		json.NewEncoder(w).Encode(badResponse)
		return
	}

	db, servicef := InitDBHandler(w)
	if db == nil {
		return
	}

	err = servicef.CreateTaskService(&task)
	if err != nil {
		badResponse := model.ResponseError{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
		}
		json.NewEncoder(w).Encode(badResponse)
		return
	}

	msg := fmt.Sprintln("Berhasil Membuat Task")
	utils.SuccesMessage(msg)

	response := model.Response{
		StatusCode: http.StatusCreated,
		Message:    "Created task success",
		Data:       task,
	}
	json.NewEncoder(w).Encode(response)
}

// curl -X PUT http://localhost:8080/api/todos/update -H "token:token_admin_1" -H "Content-Type: application/json" -d '{"id":18, "user_id":2, "title":"project golang endpoint", "description":"membuat api endpoint todo list"}'
func UpdateTaskHandler(w http.ResponseWriter, r *http.Request) {
	task := model.Task{}
	err := json.NewDecoder(r.Body).Decode(&task)

	if err != nil {
		badResponse := model.ResponseError{
			StatusCode: http.StatusBadRequest,
			Message:    "Error server",
		}
		json.NewEncoder(w).Encode(badResponse)
		return
	}

	db, servicef := InitDBHandler(w)
	if db == nil {
		return
	}

	err = servicef.UpdateTaskService(&task)
	if err != nil {
		badResponse := model.ResponseError{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
		}
		json.NewEncoder(w).Encode(badResponse)
		return
	}

	msg := fmt.Sprintln("Berhasil Memperbaharui Task")
	utils.SuccesMessage(msg)

	response := model.Response{
		StatusCode: http.StatusCreated,
		Message:    "Update task success",
		Data:       task,
	}
	json.NewEncoder(w).Encode(response)
}

// curl -X GET http://localhost:8080/api/todos/read -H "token:token_admin_1" -H "Content-Type: application/json"
func ReadTaskHandler(w http.ResponseWriter, r *http.Request) {
	task := []model.Task{}
	db, servicef := InitDBHandler(w)
	if db == nil {
		return
	}

	err := servicef.ReadTaskService(&task)
	if err != nil {
		badResponse := model.ResponseError{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
		}
		json.NewEncoder(w).Encode(badResponse)
		return
	}

	msg := fmt.Sprintln("Berhasil Membaca Task")
	utils.SuccesMessage(msg)

	response := model.Response{
		StatusCode: http.StatusCreated,
		Message:    "Get task success",
		Data:       task,
	}
	json.NewEncoder(w).Encode(response)
}

// curl -X DELETE http://localhost:8080/api/todos/delete -H "token:token_admin_1" -H "Content-Type: application/json" -d '{"id":18}'
func DeleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	task := model.Task{}
	err := json.NewDecoder(r.Body).Decode(&task)

	if err != nil {
		badResponse := model.ResponseError{
			StatusCode: http.StatusBadRequest,
			Message:    "Error server",
		}
		json.NewEncoder(w).Encode(badResponse)
		return
	}

	db, servicef := InitDBHandler(w)
	if db == nil {
		return
	}

	err = servicef.DeleteTaskService(&task)
	if err != nil {
		badResponse := model.ResponseError{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
		}
		json.NewEncoder(w).Encode(badResponse)
		return
	}

	msg := fmt.Sprintln("Berhasil Menghapus Task")
	utils.SuccesMessage(msg)

	response := model.Response{
		StatusCode: http.StatusCreated,
		Message:    "Delete task success",
		Data:       task,
	}
	json.NewEncoder(w).Encode(response)
}
