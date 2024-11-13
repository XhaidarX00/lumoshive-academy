package router

import (
	"latihan/controller"
	middlewaree "latihan/middleware"
	"log"
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

func InitRoute(ch *controller.Controller, mid *middlewaree.Middlewaree) {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	fs := http.FileServer(http.Dir("asset/"))
	r.Handle("/img/*", http.StripPrefix("/img", fs))

	// r.With(logger.MiddlewareeLogger).Route("/api/payments", func(r chi.Router) {
	// 	r.Get("/{id}", ch.Payments.GetPaymentByID)
	// r.Post("/", payments.CreatePayment)
	// r.Get("/", payments.GetAllPayments)
	// r.Get("/{id}", payments.GetPaymentByID)
	// r.Put("/{id}", payments.UpdatePayment)
	// r.Delete("/{id}", payments.DeletePayment)
	// })

	r.With(mid.MiddlewareeLogger).Route("/", func(r chi.Router) {
		r.Get("/", ch.Usershandler.LoginHandler)
		r.Get("/login", ch.Usershandler.LoginHandler)
		r.Post("/login", ch.Usershandler.LoginHandler)
		r.With(mid.TokenMiddleware).Route("/api", func(r chi.Router) {
			r.Get("/dashboard", ch.Books.DashboardHandler)
			r.Get("/book-list", ch.Books.BookListHandler)
			r.Get("/add-book", ch.Books.AddBookHandler)
			r.Get("/review", ch.Books.SubmitReviewHandler)
			r.Post("/add-book", ch.Books.AddBookHandler)

			r.Route("/book", func(r chi.Router) {
				r.Get("/discount", ch.Books.DiscountBookHandler)
				r.Route("/{bookID}", func(r chi.Router) {
					r.Get("/edit", ch.Books.EditBookHandler)
					r.Post("/edit", ch.Books.EditBookHandler)
					r.Get("/delete", ch.Books.DeleteBookHandler)
				})
			})

			r.Route("/order", func(r chi.Router) {
				r.Get("/", ch.Orders.OrderListHandler)
				r.Get("/add", ch.Orders.AddOrderHandler)
				r.Route("/{orderID}", func(r chi.Router) {
					r.Get("/", ch.Orders.OrderDetailHandler)
				})
			})

			r.Route("/payments", func(r chi.Router) {
				r.Post("/", ch.Payments.CreatePayment)
				r.Get("/", ch.Payments.GetAllPayments)
				r.Get("/{id}", ch.Payments.GetPaymentByID)
				r.Put("/{id}", ch.Payments.UpdatePayment)
				r.Delete("/{id}", ch.Payments.DeletePayment)
			})

			r.Route("/user", func(r chi.Router) {
				r.Get("/logout", ch.Usershandler.LogoutHandler)
			})
		})
	})

	r.MethodNotAllowed(MethodNotAllowedHandler)

	log.Println("server started on http://localhost:8080")
	http.ListenAndServe(":8080", r)
}

func MethodNotAllowedHandler(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
}
