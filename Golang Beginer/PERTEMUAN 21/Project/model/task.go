package model

type Task struct {
	ID          int    `json:"id"`
	User_id     int    `json:"user_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

// curl -X POST http://localhost:8080/todo/update -H "token:e675ef77-d1f9-40ab-b543-cbaf0ebbfc2d" -H "Content-Type: application/json" -d '{"id":1, "title":"project golang", "description":"membuat api endpoint todo list"}'
