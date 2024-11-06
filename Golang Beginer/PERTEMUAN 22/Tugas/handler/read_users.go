package handler

import (
	"html/template"
	"main/model"
	"main/service"
	"net/http"
)

func ReadUsers(w http.ResponseWriter, r *http.Request) {
	if token == nil {
		temp, _ := template.ParseFiles("templates/base.html", "templates/error.html")
		temp.Execute(w, map[string]string{
			"ErrorMessage": "Anda belum registrasi. Silakan melakukan registrasi terlebih dahulu.",
		})

		return
	}

	var users []model.Users
	service.ServiceF.GetUsersDataService(&users)

	data := map[string]any{
		"users": users,
	}

	temp, err := template.ParseFiles("templates/base.html", "templates/list-users.html")
	if err != nil {
		panic(err)
	}

	temp.Execute(w, data)
}
