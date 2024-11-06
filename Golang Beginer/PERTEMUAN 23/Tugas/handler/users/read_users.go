package users

import (
	"main/handler"
	UserModel "main/model/users"
	"main/service"
	"net/http"
)

func ReadUsers(w http.ResponseWriter, r *http.Request) {
	var users []UserModel.Users
	service.ServiceF.GetUsersDataService(&users)

	data := map[string]any{
		"users": users,
	}

	handler.RenderTemplate(w, "list-users.html", data)
}
