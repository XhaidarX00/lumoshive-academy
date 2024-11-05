package main

import (
	"html/template"
	"log"
	"net/http"
)

type PageData struct {
	Messages []string
}

var (
	tmpl     *template.Template
	messages []string
)

func handler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		r.ParseForm()
		message := r.FormValue("message")
		messages = append(messages, message)
	}

	data := PageData{
		Messages: messages,
	}

	err := tmpl.ExecuteTemplate(w, "index.html", data)
	if err != nil {
		log.Fatalf("Error executing template: %v", err)
	}
}

func main() {
	var err error
	tmpl, err = template.ParseFiles("index.html")
	if err != nil {
		log.Fatalf("Error parsing template: %v", err)
	}

	http.HandleFunc("/", handler)
	log.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
