package users

import (
	"html/template"
	"log"
	UserModel "main/model/users"
	"main/service"
	"net/http"
)

var Token *string

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		temp, err := template.ParseFiles("view/base.html", "view/login.html")
		if err != nil {
			panic(err)
		}

		temp.Execute(w, nil)
	}

	if r.Method == "POST" {
		var user = UserModel.Users{
			Username: r.FormValue("username"),
			Password: r.FormValue("password"),
		}
		err := service.ServiceF.LoginService(&user)
		if err != nil {
			temp, _ := template.ParseFiles("view/base.html", "view/login.html")

			temp.Execute(w, nil)
		}

		http.Redirect(w, r, "/list-users", http.StatusSeeOther)

		Token = &user.Token

		log.Println(Token)
	}
}
