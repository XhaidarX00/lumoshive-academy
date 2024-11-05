package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	// Mendapatkan informasi dari HTTP request
	method := r.Method
	url := r.URL.Path
	host := r.Host
	userAgent := r.UserAgent()
	// Menampilkan informasi request ke konsol
	fmt.Printf("HTTP Method: %s\n", method)
	fmt.Printf("URL Path: %s\n", url)
	fmt.Printf("Host: %s\n", host)
	fmt.Printf("User-Agent: %s\n", userAgent)

	// Mengirimkan respons ke client
	fmt.Fprintf(w, "Hello, World!")
}

func main() {
	// Mendefinisikan handler untuk route "/"
	http.HandleFunc("/", helloHandler)

	// Menjalankan server HTTP pada port 8080
	fmt.Println("Server started on port 8080")
	http.ListenAndServe(":8080", nil)
}
