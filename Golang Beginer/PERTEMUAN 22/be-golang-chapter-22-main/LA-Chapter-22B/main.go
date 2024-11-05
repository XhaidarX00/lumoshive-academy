// // example template data
// package main

// import (
// 	"html/template"
// 	"log"
// 	"net/http"
// )

// // Struct untuk data yang akan dikirim ke template
// type PageData struct {
// 	Title   string
// 	Name    string
// 	Message string
// }

// func handler(w http.ResponseWriter, r *http.Request) {
// 	// Parse template
// 	tmpl, err := template.ParseFiles("template.html")
// 	if err != nil {
// 		log.Fatalf("Error parsing template: %v", err)
// 	}

// 	// Data yang akan dikirim ke template
// 	data := PageData{
// 		Title:   "Welcome Page",
// 		Name:    "Lumoshive",
// 		Message: "Welcome to the Go HTML template example.",
// 	}

// 	// Render template dengan data
// 	err = tmpl.Execute(w, data)
// 	if err != nil {
// 		log.Fatalf("Error executing template: %v", err)
// 	}
// }

// func main() {
// 	http.HandleFunc("/", handler)
// 	log.Println("Server running on http://localhost:8080")
// 	log.Fatal(http.ListenAndServe(":8080", nil))
// }

// // example template action
// package main

// import (
// 	"html/template"
// 	"log"
// 	"net/http"
// )

// type Person struct {
// 	Nama string
// 	Umur int
// }

// func handler(w http.ResponseWriter, r *http.Request) {
// 	// Contoh data
// 	data := Person{
// 		Nama: "Ani",
// 		Umur: 15,
// 	}

// 	tmpl, err := template.ParseFiles("template.html")
// 	if err != nil {
// 		log.Fatalf("Error parsing template: %v", err)
// 	}

// 	err = tmpl.Execute(w, data)
// 	if err != nil {
// 		log.Fatalf("Error executing template: %v", err)
// 	}
// }

// func main() {
// 	http.HandleFunc("/", handler)
// 	log.Println("Server running on http://localhost:8080")
// 	log.Fatal(http.ListenAndServe(":8080", nil))
// }

// // example operator perbandingan
// package main

// import (
// 	"html/template"
// 	"log"
// 	"net/http"
// )

// type Person struct {
// 	Nama string
// 	Umur int
// }

// func handler(w http.ResponseWriter, r *http.Request) {
// 	// Contoh data
// 	data := Person{
// 		Nama: "Budi",
// 		Umur: 20,
// 	}

// 	tmpl, err := template.ParseFiles("template.html")
// 	if err != nil {
// 		log.Fatalf("Error parsing template: %v", err)
// 	}

// 	err = tmpl.Execute(w, data)
// 	if err != nil {
// 		log.Fatalf("Error executing template: %v", err)
// 	}
// }

// func main() {
// 	http.HandleFunc("/", handler)
// 	log.Println("Server running on http://localhost:8080")
// 	log.Fatal(http.ListenAndServe(":8080", nil))
// }

// // example  action range
// package main

// import (
// 	"html/template"
// 	"log"
// 	"net/http"
// )

// func handler(w http.ResponseWriter, r *http.Request) {
// 	// Contoh data
// 	data := struct {
// 		Buah []string
// 	}{
// 		Buah: []string{"Apel", "Jeruk", "Mangga", "Anggur"},
// 	}
// 	tmpl, err := template.ParseFiles("template.html")
// 	if err != nil {
// 		log.Fatalf("Error parsing template: %v", err)
// 	}

// 	err = tmpl.Execute(w, data)
// 	if err != nil {
// 		log.Fatalf("Error executing template: %v", err)
// 	}
// }

// func main() {
// 	http.HandleFunc("/", handler)
// 	log.Println("Server running on http://localhost:8080")
// 	log.Fatal(http.ListenAndServe(":8080", nil))
// }

// example action with
package main

import (
	"html/template"
	"log"
	"net/http"
)

type Address struct {
	Street  string
	City    string
	Country string
}
type Person struct {
	Name    string
	Age     int
	Address Address
}

func handler(w http.ResponseWriter, r *http.Request) {
	// Contoh data dengan nested struct
	data := Person{
		Name: "John Doe",
		Age:  30,
		Address: Address{
			Street:  "Jl. Raya No. 123",
			City:    "Jakarta",
			Country: "Indonesia",
		},
	}
	tmpl, err := template.ParseFiles("template.html")
	if err != nil {
		log.Fatalf("Error parsing template: %v", err)
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		log.Fatalf("Error executing template: %v", err)
	}
}

func main() {
	http.HandleFunc("/", handler)
	log.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
