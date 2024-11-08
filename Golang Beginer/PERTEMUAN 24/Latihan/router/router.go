package router

import (
	bookstore "latihan/controller/bookStore"
	"latihan/controller/orders"
	usershandler "latihan/controller/users_handler"
	middlewaree "latihan/middleware"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func InitRoute() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Route("/", func(r chi.Router) {
		r.Get("/", usershandler.LoginHandler)
		r.Get("/login", usershandler.LoginHandler)
		r.Post("/login", usershandler.LoginHandler)

		r.With(middlewaree.TokenMiddleware).Route("/api", func(r chi.Router) {
			r.Get("/dashboard", bookstore.DashboardHandler)
			r.Get("/book-list", bookstore.BookListHandler)
			r.Get("/add-book", bookstore.AddBookHandler)
			r.Post("/add-book", bookstore.AddBookHandler)

			r.Route("/user", func(r chi.Router) {
				r.Get("/logout", usershandler.LogoutHandler)
			})

			r.Route("/book", func(r chi.Router) {
				r.Get("/discount", bookstore.DiscountBookHandler)
				r.Route("/{bookID}", func(r chi.Router) {
					r.Get("/edit", bookstore.EditBookHandler)
					r.Post("/edit", bookstore.EditBookHandler)
					r.Get("/delete", bookstore.DeleteBookHandler)
				})
			})

			r.Route("/order", func(r chi.Router) {
				r.Get("/", orders.OrderListHandler)
				r.Route("/{orderID}", func(r chi.Router) {
					r.Get("/detail", orders.OrderDetailHandler)
				})
			})
		})
	})

	r.MethodNotAllowed(MethodNotAllowedHandler)

	log.Println("server started on http://localhost:8000")
	http.ListenAndServe(":8000", r)
}

func MethodNotAllowedHandler(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
}
