package handler

import (
	"main/service"
	"net/http"
	"strconv"
)

func DeleteUsers(w http.ResponseWriter, r *http.Request) {
	idString := r.URL.Query().Get("id")

	id, err := strconv.Atoi(idString)
	if err != nil {
		panic(err)
	}

	if err := service.ServiceF.DeleteUserService(id); err != nil {
		panic(err)
	}

	http.Redirect(w, r, "/list-users", http.StatusSeeOther)
}
