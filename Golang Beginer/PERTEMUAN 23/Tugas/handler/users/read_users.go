package users

import (
	"html/template"

	UserModel "main/model/users"
	"main/service"
	"net/http"
)

func ReadUsers(w http.ResponseWriter, r *http.Request) {
	if Token == nil {
		temp, _ := template.ParseFiles("view/base.html", "view/error.html")
		temp.Execute(w, map[string]string{
			"ErrorMessage": "Anda belum registrasi. Silakan melakukan registrasi terlebih dahulu.",
		})

		return
	}

	var users []UserModel.Users
	service.ServiceF.GetUsersDataService(&users)

	data := map[string]any{
		"users": users,
	}

	temp, err := template.ParseFiles("view/base.html", "view/list-users.html")
	if err != nil {
		panic(err)
	}

	temp.Execute(w, data)
}
