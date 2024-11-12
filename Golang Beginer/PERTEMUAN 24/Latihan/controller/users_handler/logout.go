package usershandler

import (
	pagehandler "latihan/controller/pageHandler"
	"net/http"
)

func (a *Auth) LogoutHandler(w http.ResponseWriter, r *http.Request) {
	pagehandler.RenderTemplate(w, "logout.html", nil)
}
