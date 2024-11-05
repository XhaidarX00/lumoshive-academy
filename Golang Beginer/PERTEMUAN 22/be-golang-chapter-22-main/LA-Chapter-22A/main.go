package main

import (
	"html/template"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("index.html")
	if err != nil {
		log.Fatalf("Error parsing template: %v", err)
	}

	err = tmpl.Execute(w, "lumoshive academy")
	if err != nil {
		log.Fatalf("Error executing template: %v", err)
	}
}

func main() {
	http.HandleFunc("/", handler)
	log.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
