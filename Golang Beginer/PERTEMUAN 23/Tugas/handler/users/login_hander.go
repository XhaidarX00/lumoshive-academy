package users

import (
	"main/handler"
	"main/library"
	UserModel "main/model/users"
	"main/service"
	"net/http"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		handler.RenderTemplate(w, "login.html", nil)
	}

	if r.Method == "POST" {
		var user = UserModel.Users{
			Username: r.FormValue("username"),
			Password: r.FormValue("password"),
		}
		err := service.ServiceF.LoginService(&user)
		if err != nil {
			handler.RenderTemplate(w, "login.html", nil)
		}

		library.SetCookie(w, user.Token)
		http.Redirect(w, r, "/api/user", http.StatusSeeOther)
	}
}
