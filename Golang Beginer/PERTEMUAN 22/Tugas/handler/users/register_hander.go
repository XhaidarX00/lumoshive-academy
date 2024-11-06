package users

import (
	"html/template"
	"log"
	UserModel "main/model/users"
	"main/service"
	"net/http"
)

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		temp, err := template.ParseFiles("view/base.html", "view/registration.html")
		if err != nil {
			panic(err)
		}

		temp.Execute(w, nil)
	}

	if r.Method == "POST" {
		var user = UserModel.Users{
			Name:     r.FormValue("name"),
			Username: r.FormValue("username"),
			Password: r.FormValue("password"),
		}
		err := service.ServiceF.RegisterService(&user)
		if err != nil {
			temp, _ := template.ParseFiles("view/base.html", "view/registration.html")

			temp.Execute(w, nil)
		}

		http.Redirect(w, r, "/login", http.StatusSeeOther)

		log.Printf("Register Berhasil : %v\n", user)
	}

}
