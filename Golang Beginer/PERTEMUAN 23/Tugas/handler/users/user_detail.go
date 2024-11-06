package users

import (
	"html/template"

	UserModel "main/model/users"
	"main/service"
	"net/http"
	"strconv"
)

func UserDetail(w http.ResponseWriter, r *http.Request) {
	var users UserModel.Users
	idStr := r.URL.Query().Get("id")
	idInt, err := strconv.Atoi(idStr)
	if err != nil {
		return
	}
	users.ID = idInt
	service.ServiceF.GetUsersDetailService(&users)

	data := map[string]any{
		"user": users,
	}

	temp, err := template.ParseFiles("view/base.html", "view/user-detail.html")
	if err != nil {
		panic(err)
	}

	temp.Execute(w, data)

}
