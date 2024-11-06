package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func mainsdsa() {
	r := chi.NewRouter()

	r.Route("/api", func(r chi.Router) {

		// Grouping /api/users routes
		r.Route("/users", func(r chi.Router) {
			r.Get("/", getAllUsers)
			r.Post("/", createUser)
			r.Route("/{userID}", func(r chi.Router) {
				r.Get("/", getUserByID)
				r.Put("/", updateUser)
				r.Delete("/", deleteUser)
			})
		})

		// Grouping /api/products routes
		r.Route("/products", func(r chi.Router) {
			r.Get("/", getAllProducts)
			r.Post("/", createProduct)
			r.Route("/{productID}", func(r chi.Router) {
				r.Get("/", getProductByID)
				r.Put("/", updateProduct)
				r.Delete("/", deleteProduct)
			})
		})
	})

	http.ListenAndServe(":8080", r)
}

// Handler functions for users
func getAllUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Get all users"))
}

func createUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create a new user"))
}

func getUserByID(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "userID")
	w.Write([]byte("Get user by ID: " + userID))
}

func updateUser(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "userID")
	w.Write([]byte("Update user by ID: " + userID))
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "userID")
	w.Write([]byte("Delete user by ID: " + userID))
}

// Handler functions for products
func getAllProducts(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Get all products"))
}

func createProduct(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create a new product"))
}

func getProductByID(w http.ResponseWriter, r *http.Request) {
	productID := chi.URLParam(r, "productID")
	w.Write([]byte("Get product by ID: " + productID))
}

func updateProduct(w http.ResponseWriter, r *http.Request) {
	productID := chi.URLParam(r, "productID")
	w.Write([]byte("Update product by ID: " + productID))
}

func deleteProduct(w http.ResponseWriter, r *http.Request) {
	productID := chi.URLParam(r, "productID")
	w.Write([]byte("Delete product by ID: " + productID))
}
