package handler

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

// Template cache untuk menyimpan template yang sudah diparsing
var templates = map[string]*template.Template{}

// InitTemplates - Memuat semua template saat aplikasi dimulai
func InitTemplates() {
	baseTemplate := "view/base.html"
	contentTemplates := []string{
		"view/list-users.html",
		"view/registration.html",
		"view/todo-list.html",
		"view/user-detail.html",
		"view/error.html",
		"view/login.html",
	}

	for _, content := range contentTemplates {
		name := filepath.Base(content)
		tmpl, err := template.ParseFiles(baseTemplate, content)
		if err != nil {
			log.Fatalf("Error parsing templates: %v", err)
		}
		templates[name] = tmpl
	}
}

// RenderTemplate - Merender template berdasarkan nama file
func RenderTemplate(w http.ResponseWriter, tmplName string, data interface{}) {
	tmpl, ok := templates[tmplName]
	if !ok {
		http.Error(w, "Template not found", http.StatusNotFound)
		return
	}

	err := tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		log.Println(err)
	}
}

// Handler untuk setiap halaman
// func RegistrationHandler(w http.ResponseWriter, r *http.Request) {
// 	RenderTemplate(w, "registration.html", nil)
// }

// var templates = template.Must(template.ParseGlob("view/*.html"))
