package todos

import (
	"fmt"
	"log"
	"main/handler"
	todosModel "main/model/todos"
	"main/service"
	"net/http"
)

func AddTodo(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		http.Redirect(w, r, "/api/todos", http.StatusSeeOther)
		return
	}

	if r.Method == "POST" {
		var task = todosModel.Task{
			Title:       r.FormValue("title"),
			Description: fmt.Sprintf("Deskripsi %s\n", r.FormValue("title")),
		}

		id, err := service.ServiceF.AddTodoService(&task)
		if err != nil {
			handler.RenderTemplate(w, "todo-list.html", nil)
			return
		}

		log.Printf("Berhasil membuat task dengan id : %d\n", id)

		http.Redirect(w, r, "/api/todos/", http.StatusSeeOther)
	}
}
