// // example upload file
// package main

// import (
// 	"fmt"
// 	"html/template"
// 	"io"
// 	"log"
// 	"net/http"
// 	"os"
// 	"path/filepath"
// )

// func uploadFile(w http.ResponseWriter, r *http.Request) {
// 	// Batasi ukuran file yang diunggah, misalnya 10 MB
// 	r.ParseMultipartForm(10 << 20)

// 	// Ambil file dari form
// 	file, handler, err := r.FormFile("file")
// 	if err != nil {
// 		fmt.Println("Error Retrieving the File")
// 		fmt.Println(err)
// 		return
// 	}
// 	defer file.Close()
// 	fmt.Printf("Uploaded File: %+v\n", handler.Filename)
// 	fmt.Printf("File Size: %+v\n", handler.Size)
// 	fmt.Printf("MIME Header: %+v\n", handler.Header)

// 	// Buat file di server
// 	dst, err := os.Create(filepath.Join("uploads", handler.Filename))
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	defer dst.Close()

// 	// Salin konten file ke tujuan
// 	_, err = io.Copy(dst, file)

// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}

// 	fmt.Fprintf(w, "Successfully Uploaded File\n")
// }

// func FormUpload(w http.ResponseWriter, r *http.Request) {
// 	data := struct {
// 		Title string
// 	}{
// 		Title: "lumoshive",
// 	}

// 	// Parse layout template
// 	tmpl, err := template.ParseFiles("index.html")
// 	if err != nil {
// 		log.Fatalf("Error parsing template: %v", err)
// 	}

// 	// Execute layout template
// 	err = tmpl.Execute(w, data)
// 	if err != nil {
// 		log.Fatalf("Error executing template: %v", err)
// 	}

// }

// func main() {
// 	http.HandleFunc("/upload", uploadFile)
// 	http.HandleFunc("/", FormUpload)
// 	log.Fatal(http.ListenAndServe(":8080", nil))
// }

// example download file
package main

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
)

func downloadFile(w http.ResponseWriter, r *http.Request) {
	// Path ke file yang akan diunduh
	filePath := "./uploads/sample.pdf"

	// Buka file
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Set header untuk response
	w.Header().Set("Content-Disposition", "attachment; filename="+filePath)
	w.Header().Set("Content-Type", "application/pdf")
	w.Header().Set("Content-Length", fmt.Sprintf("%d", fileSize(filePath)))

	// Salin file ke response writer
	_, err = io.Copy(w, file)

	if err != nil {
		log.Fatal(err)
	}
}

func fileSize(filePath string) int64 {
	fi, err := os.Stat(filePath)
	if err != nil {
		log.Fatal(err)
	}
	return fi.Size()
}

func formDonload(w http.ResponseWriter, r *http.Request) {
	data := struct {
		Title string
	}{
		Title: "Lumoshive",
	}

	// Parse layout template
	tmpl, err := template.ParseFiles("index.html")
	if err != nil {
		log.Fatalf("Error parsing template: %v", err)
	}

	// Execute layout template
	err = tmpl.Execute(w, data)
	if err != nil {
		log.Fatalf("Error executing template: %v", err)
	}

}

func main() {
	http.HandleFunc("/download", downloadFile)
	http.HandleFunc("/", formDonload)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
