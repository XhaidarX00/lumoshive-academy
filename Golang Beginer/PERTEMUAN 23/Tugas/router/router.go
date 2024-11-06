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
		r.Get("/register", users.RegisterHandler)
		r.Post("/register", users.RegisterHandler)
		r.Get("/login", users.LoginHandler)
		r.Post("/login", users.LoginHandler)

		r.With(middlewaree.TokenMiddleware).Route("/api", func(r chi.Router) {
			r.Route("/users", func(r chi.Router) {
				r.Get("/", users.ReadUsers)
				r.Route("/{userID}", func(r chi.Router) {
					r.Get("/", users.UserDetail)
					r.Delete("/", users.DeleteUsers)
				})
			})

			r.Route("/todos", func(r chi.Router) {
				r.Get("/", todos.ReadTodos)
				r.Post("/", todos.AddTodo)
				r.Route("/{taskID}", func(r chi.Router) {
					r.Put("/", todos.ChangeTodoStatus)
					r.Delete("/", todos.DeleteTodo)
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
