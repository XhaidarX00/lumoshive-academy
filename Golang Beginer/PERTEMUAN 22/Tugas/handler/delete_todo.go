package handler

import (
	"log"
	"main/service"
	"net/http"
	"strconv"
)

func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	idString := r.URL.Query().Get("id")

	id, err := strconv.Atoi(idString)
	if err != nil {
		panic(err)
	}

	if err := service.ServiceF.DeleteTodoService(id); err != nil {
		panic(err)
	}

	log.Printf("Berhasil menghapus todo id %d\n", id)
	http.Redirect(w, r, "/todo-list", http.StatusSeeOther)
}
