package users

import (
	"main/service"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

func DeleteUsers(w http.ResponseWriter, r *http.Request) {
	idString := chi.URLParam(r, "userID")

	id, err := strconv.Atoi(idString)
	if err != nil {
		panic(err)
	}

	if err := service.ServiceF.DeleteUserService(id); err != nil {
		panic(err)
	}

	http.Redirect(w, r, "/api/user/", http.StatusSeeOther)
}
