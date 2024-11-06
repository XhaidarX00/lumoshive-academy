package todos

import (
	"log"
	"main/service"
	"net/http"
	"strconv"
)

func ChangeTodoStatus(w http.ResponseWriter, r *http.Request) {
	idString := r.URL.Query().Get("id")

	id, err := strconv.Atoi(idString)
	if err != nil {
		panic(err)
	}

	if err := service.ServiceF.ChangeTodoStatusService(id); err != nil {
		panic(err)
	}

	log.Printf("Berhasil Update Status %d\n", id)

	http.Redirect(w, r, "/todo-list", http.StatusSeeOther)
}
