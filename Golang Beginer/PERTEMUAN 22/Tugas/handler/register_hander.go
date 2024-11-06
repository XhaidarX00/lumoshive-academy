package handler

import (
	"html/template"
	"log"
	"main/model"
	"main/service"
	"net/http"
)

var token *string

// curl -X POST http://localhost:8080/api/auth/register -H "Content-Type: application/json" -d '{"username":"admin3", "password":"password789", "email":"admin3@example.com", "role":"admin"}'
func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		temp, err := template.ParseFiles("templates/base.html", "templates/registration.html")
		if err != nil {
			panic(err)
		}

		temp.Execute(w, nil)
	}

	if r.Method == "POST" {
		var user = model.Users{
			Name:     r.FormValue("name"),
			Username: r.FormValue("username"),
			Password: r.FormValue("password"),
		}
		err := service.ServiceF.RegisterService(&user)
		if err != nil {
			temp, _ := template.ParseFiles("templates/base.html", "templates/registration.html")

			temp.Execute(w, nil)
		}

		http.Redirect(w, r, "/list-users", http.StatusSeeOther)

		log.Printf("Register Berhasil : %v\n", user)

		token = &user.Token

		log.Println(token)
	}

	// fmt.Println("Masuk Register")
	// user := model.Users{}
	// err := json.NewDecoder(r.Body).Decode(&user)
	// if err != nil {
	// 	library.ResponseToJson(w, "Error Server", nil)
	// 	return
	// }

	// if len(user.Password) < 8 {
	// 	library.ResponseToJson(w, "Must be at lease 8 character in password", nil)
	// 	return
	// }

	// err = service.ServiceF.RegisterService(&user)
	// if err != nil {
	// 	library.ResponseToJson(w, err.Error(), nil)
	// 	return
	// }

	// response := library.CreatedRequest("Register success", user)
	// library.JsonResponse(w, response)
}
