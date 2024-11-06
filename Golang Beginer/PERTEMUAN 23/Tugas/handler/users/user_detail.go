package users

import (
	"main/handler"
	UserModel "main/model/users"
	"main/service"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

func UserDetail(w http.ResponseWriter, r *http.Request) {
	var users UserModel.Users
	idStr := chi.URLParam(r, "userID")
	idInt, err := strconv.Atoi(idStr)
	if err != nil {
		return
	}
	users.ID = idInt
	service.ServiceF.GetUsersDetailService(&users)

	data := map[string]any{
		"user": users,
	}

	handler.RenderTemplate(w, "user-detail.html", data)
}
