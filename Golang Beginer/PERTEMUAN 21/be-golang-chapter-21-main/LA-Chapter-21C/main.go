package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type requestData struct {
	Name  string
	Age   int
	Email string
}

func queryHandler(w http.ResponseWriter, r *http.Request) {
	// Mengambil query parameter dari URL
	query := r.URL.Query()

	// Mendapatkan nilai dari query parameter tertentu
	name := query.Get("name")
	age := query.Get("age")

	// Memeriksa dan menggunakan nilai query parameter
	if name == "" {
		name = "Guest"
	}

	if age == "" {
		age = "unknown"
	}

	requestData := requestData{}

	// Decode JSON dari body request ke struct
	err := json.NewDecoder(r.Body).Decode(&requestData)
	if err != nil {
		http.Error(w, "Failed to parse JSON", http.StatusBadRequest)
		return
	}

	// Contoh penggunaan data yang diambil
	fmt.Fprintf(w, "Received Name: %s, Age: %d, Email: %s", requestData.Name, requestData.Age, requestData.Email)

	// Mengirimkan respons ke client
	fmt.Fprintf(w, "Hello, %s! Your age is %s.\n", name, age)
}

func handler(w http.ResponseWriter, r *http.Request) {
	// Membaca header dari request
	userAgent := r.Header.Get("User-Agent")
	authorization := r.Header.Get("Authorization")
	// Menampilkan header di server log
	fmt.Printf("User-Agent: %s\n", userAgent)
	fmt.Printf("Authorization: %s\n", authorization)

	// Mengirimkan respons
	fmt.Fprintf(w, `{"message": "Headers received"}`)
}

func main() {
	// Mendaftarkan handler function untuk route "/query"
	http.HandleFunc("/query", queryHandler)

	http.HandleFunc("/", handler)
	// Menjalankan server HTTP pada port 8080
	fmt.Println("Server started at :8080")
	http.ListenAndServe(":8080", nil)
}
