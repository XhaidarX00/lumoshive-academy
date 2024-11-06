package handler

import (
	"html/template"
	"main/model"
	"main/service"
	"net/http"
)

func ReadTodos(w http.ResponseWriter, r *http.Request) {
	if token == nil {
		temp, _ := template.ParseFiles("templates/base.html", "templates/error.html")
		temp.Execute(w, map[string]string{
			"ErrorMessage": "Anda belum registrasi. Silakan melakukan registrasi terlebih dahulu.",
		})

		return
	}

	var tasks []model.Task
	service.ServiceF.GetTodosService(&tasks)

	data := map[string]any{
		"todos": tasks,
	}

	temp, err := template.ParseFiles("templates/base.html", "templates/todo-list.html")
	if err != nil {
		panic(err)
	}

	temp.Execute(w, data)
}
