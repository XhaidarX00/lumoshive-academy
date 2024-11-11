package router

import (
	payments "latihan/controller/payment"
	"log"
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

func InitRoute() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	fs := http.FileServer(http.Dir("asset/"))
	r.Handle("/img/*", http.StripPrefix("/img", fs))

	r.Route("/api/payments", func(r chi.Router) {
		r.Post("/", payments.CreatePayment)
		r.Get("/", payments.GetAllPayments)
		r.Get("/{id}", payments.GetPaymentByID)
		r.Put("/{id}", payments.UpdatePayment)
		r.Delete("/{id}", payments.DeletePayment)
	})

	// r.Route("/", func(r chi.Router) {

	// r.Get("/", usershandler.LoginHandler)
	// r.Get("/login", usershandler.LoginHandler)
	// r.Post("/login", usershandler.LoginHandler)
	// r.With(middlewaree.TokenMiddleware).Route("/api", func(r chi.Router) {
	// r.Get("/dashboard", bookstore.DashboardHandler)
	// r.Get("/book-list", bookstore.BookListHandler)
	// r.Get("/add-book", bookstore.AddBookHandler)
	// r.Post("/add-book", bookstore.AddBookHandler)

	// r.Route("/user", func(r chi.Router) {
	// 	r.Get("/logout", usershandler.LogoutHandler)
	// })

	// r.Route("/book", func(r chi.Router) {
	// 	r.Get("/discount", bookstore.DiscountBookHandler)
	// 	r.Route("/{bookID}", func(r chi.Router) {
	// 		r.Get("/edit", bookstore.EditBookHandler)
	// 		r.Post("/edit", bookstore.EditBookHandler)
	// 		r.Get("/delete", bookstore.DeleteBookHandler)
	// 	})
	// })

	// r.Route("/order", func(r chi.Router) {
	// 	r.Get("/", orders.OrderListHandler)
	// 	r.Route("/{orderID}", func(r chi.Router) {
	// 		r.Get("/detail", orders.OrderDetailHandler)
	// 	})
	// })

	// r.Route("/payments", func(r chi.Router) {
	// 	r.Post("/", paymentHandler.CreatePayment)
	// 	r.Get("/", paymentHandler.GetAllPayments)
	// 	r.Get("/{id}", paymentHandler.GetPaymentByID)
	// 	r.Put("/{id}", paymentHandler.UpdatePayment)
	// 	r.Delete("/{id}", paymentHandler.DeletePayment)
	// })
	// })
	// })

	r.MethodNotAllowed(MethodNotAllowedHandler)

	log.Println("server started on http://localhost:8080")
	http.ListenAndServe(":8080", r)
}

func MethodNotAllowedHandler(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
}
