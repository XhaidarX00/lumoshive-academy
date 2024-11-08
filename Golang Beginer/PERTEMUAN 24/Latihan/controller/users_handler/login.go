package usershandler

import (
	"latihan/controller"
	"latihan/library"
	"latihan/model/customers"
	"latihan/service"
	"net/http"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		controller.RenderTemplate(w, "login.html", nil)
		return
	}

	if r.Method == "POST" {
		var customer = customers.Customer{
			Username: r.FormValue("username"),
			Password: r.FormValue("password"),
		}
		err := service.ServiceF.LoginService(&customer)
		if err != nil {
			controller.ErrorPage(w, err.Error())
			return
		}

		library.SetCookie(w, customer.Token)
		http.Redirect(w, r, "/api/dashboard", http.StatusSeeOther)
	}
}
