package router

import (
	"log"
	"main/handler/todos"
	"main/handler/users"
	middlewaree "main/middleware"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func InitRoute() {
	r := chi.NewRouter()

	r.Use(middleware.Logger)

	r.Route("/", func(r chi.Router) {
		r.Get("/", users.LoginHandler)
		r.Get("/login", users.LoginHandler)
		r.Post("/login", users.LoginHandler)
		r.Get("/registration", users.RegisterHandler)
		r.Post("/registration", users.RegisterHandler)

		r.With(middlewaree.TokenMiddleware).Route("/api", func(r chi.Router) {
			r.Route("/user", func(r chi.Router) {
				r.Get("/", users.ReadUsers)
				r.Route("/{userID}", func(r chi.Router) {
					r.Get("/", users.UserDetail)
					r.Get("/delete", users.DeleteUsers)
				})
			})

			r.Route("/todos", func(r chi.Router) {
				r.Get("/", todos.ReadTodos)
				r.Post("/", todos.AddTodo)
				r.Route("/{taskID}", func(r chi.Router) {
					r.Get("/update", todos.ChangeTodoStatus)
					r.Get("/delete", todos.DeleteTodo)
				})
			})
		})

		r.MethodNotAllowed(MethodNotAllowedHandler)
	})

	log.Println("server started on http://localhost:8080")
	http.ListenAndServe(":8080", r)
}

func MethodNotAllowedHandler(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
}
