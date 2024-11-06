package todos

import (
	"main/handler"
	todosModel "main/model/todos"
	"main/service"
	"net/http"
)

func ReadTodos(w http.ResponseWriter, r *http.Request) {
	var tasks []todosModel.Task
	service.ServiceF.GetTodosService(&tasks)

	data := map[string]any{
		"todos": tasks,
	}

	handler.RenderTemplate(w, "todo-list.html", data)
}
