package users

import (
	"log"
	"main/handler"
	UserModel "main/model/users"
	"main/service"
	"net/http"
)

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		handler.RenderTemplate(w, "registration.html", nil)
	}

	if r.Method == "POST" {
		var user = UserModel.Users{
			Name:     r.FormValue("name"),
			Username: r.FormValue("username"),
			Password: r.FormValue("password"),
		}
		err := service.ServiceF.RegisterService(&user)
		if err != nil {
			handler.RenderTemplate(w, "registration.html", nil)
		}

		http.Redirect(w, r, "/", http.StatusSeeOther)

		log.Printf("Register Berhasil : %v\n", user)
	}

}
