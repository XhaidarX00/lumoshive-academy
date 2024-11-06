package handler

import (
	"html/template"
	"main/model"
	"main/service"
	"net/http"
	"strconv"
)

func UserDetail(w http.ResponseWriter, r *http.Request) {
	var users model.Users
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

	temp, err := template.ParseFiles("templates/base.html", "templates/user-detail.html")
	if err != nil {
		panic(err)
	}

	temp.Execute(w, data)

}
