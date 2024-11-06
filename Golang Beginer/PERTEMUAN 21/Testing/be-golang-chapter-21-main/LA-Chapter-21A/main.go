// // handler
// package main

// import (
// 	"fmt"
// 	"log"
// 	"net/http"
// )

// type HelloHandler struct{}

// func (h *HelloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Hello, World!")
// }

// func main() {
// 	handler := &HelloHandler{}
// 	http.Handle("/", handler)
// 	log.Printf("Starting server on :8080")
// 	http.ListenAndServe(":8080", nil)
// }

// // function handler
// package main

// import (
// 	"fmt"
// 	"log"
// 	"net/http"
// )

// func helloHandler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Hello, World!")
// }

// func main() {
// 	http.HandleFunc("/", helloHandler)
// 	log.Printf("Starting server on :8080")
// 	http.ListenAndServe(":8080", nil)
// }

// server mux
package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to the homepage!")
	})

	mux.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "About page")
	})

	mux.HandleFunc("/contact", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Contact us")
	})

	fmt.Println("Server started on port 8080")
	http.ListenAndServe(":8080", mux)
}
