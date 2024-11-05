package handler

import "net/http"

// curl -X PUT http://localhost:8080/api/todos/update-dev -H "token:token_dev_1" -H "Content-Type: application/json" -d '{"id":20, "user_id":2, "title":"project golang endpoint", "description":"membuat api endpoint todo list"}'
func UpdateTaskHandlerDev(w http.ResponseWriter, r *http.Request) {
	UpdateTaskHandler(w, r)
}
