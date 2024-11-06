package handler

import (
	"fmt"
	"html/template"
	"log"
	"main/model"
	"main/service"
	"net/http"
)

func AddTodo(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		http.Redirect(w, r, "/todo-list", http.StatusSeeOther)
		return
	}

	if r.Method == "POST" {
		var task = model.Task{
			Title:       r.FormValue("title"),
			Description: fmt.Sprintf("Deskripsi %s\n", r.FormValue("title")),
		}

		id, err := service.ServiceF.AddTodoService(&task)
		if err != nil {
			temp, _ := template.ParseFiles("templates/base.html", "templates/todo-list.html")

			temp.Execute(w, nil)
			return
		}

		log.Printf("Berhasil membuat task dengan id : %d\n", id)

		http.Redirect(w, r, "/todo-list", http.StatusSeeOther)
	}
}
