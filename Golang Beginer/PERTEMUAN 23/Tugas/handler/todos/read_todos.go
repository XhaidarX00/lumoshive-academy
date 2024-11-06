package todos

import (
	"html/template"
	"main/handler/users"
	todosModel "main/model/todos"
	"main/service"
	"net/http"
)

func ReadTodos(w http.ResponseWriter, r *http.Request) {
	if users.Token == nil {
		temp, _ := template.ParseFiles("view/base.html", "view/error.html")
		temp.Execute(w, map[string]string{
			"ErrorMessage": "Anda belum registrasi. Silakan melakukan registrasi terlebih dahulu.",
		})

		return
	}

	var tasks []todosModel.Task
	service.ServiceF.GetTodosService(&tasks)

	data := map[string]any{
		"todos": tasks,
	}

	temp, err := template.ParseFiles("view/base.html", "view/todo-list.html")
	if err != nil {
		panic(err)
	}

	temp.Execute(w, data)
}
