package todos

import (
	"log"
	"main/service"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	idString := chi.URLParam(r, "taskID")

	id, err := strconv.Atoi(idString)
	if err != nil {
		panic(err)
	}

	if err := service.ServiceF.DeleteTodoService(id); err != nil {
		panic(err)
	}

	log.Printf("Berhasil menghapus todo id %d\n", id)
	http.Redirect(w, r, "/api/todos/", http.StatusSeeOther)
}
