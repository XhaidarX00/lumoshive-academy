package todos

import (
	"log"
	"main/service"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

func ChangeTodoStatus(w http.ResponseWriter, r *http.Request) {
	idString := chi.URLParam(r, "taskID")

	id, err := strconv.Atoi(idString)
	if err != nil {
		panic(err)
	}

	if err := service.ServiceF.ChangeTodoStatusService(id); err != nil {
		panic(err)
	}

	log.Printf("Berhasil Update Status %d\n", id)

	http.Redirect(w, r, "/api/todos/", http.StatusSeeOther)
}
