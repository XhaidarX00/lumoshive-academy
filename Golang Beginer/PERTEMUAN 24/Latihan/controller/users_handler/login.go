package usershandler

import (
	pagehandler "latihan/controller/pageHandler"
	"latihan/library"
	"latihan/model/customers"
	"latihan/service"
	"net/http"

	"go.uber.org/zap"
)

type Auth struct {
	Service *service.Service
	logger  *zap.Logger
}

func NewUserHandelr(serv *service.Service, log *zap.Logger) *Auth {
	return &Auth{
		Service: serv,
		logger:  log,
	}
}

func (a *Auth) LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		pagehandler.RenderTemplate(w, "login.html", nil)
		return
	}

	if r.Method == "POST" {
		var customer = customers.Customer{
			Username: r.FormValue("username"),
			Password: r.FormValue("password"),
		}
		err := a.Service.LoginService(&customer)
		if err != nil {
			a.logger.Error("Error LoginHandler", zap.Error(err))
			pagehandler.ErrorPage(w, err.Error())
			return
		}

		library.SetCookie(w, customer.Token)
		http.Redirect(w, r, "/api/dashboard", http.StatusSeeOther)
	}
}
