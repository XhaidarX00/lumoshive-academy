package usershandler

import (
	"latihan/controller"
	"net/http"
)

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	controller.RenderTemplate(w, "logout.html", nil)
}
