// // file server
// package main

// import (
// 	"net/http"
// )

// func main() {
// 	// Menentukan direktori yang akan disajikan
// 	fs := http.FileServer(http.Dir("./view"))

// 	// Menggunakan FileServer untuk menyajikan file dari root ("/")
// 	http.Handle("/", fs)

// 	// fs := http.FileServer(http.Dir("view"))
// 	// http.Handle("/view/", http.StripPrefix("/view/", fs))

// 	// Menjalankan server HTTP di port 8080
// 	if err := http.ListenAndServe(":8080", nil); err != nil {
// 		panic(err)
// 	}
// }

// server file
package main

import (
	"net/http"
)

func main() {
	// Mengatur route untuk melayani file statis
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {

		// Melayani file HTML ke ResponseWriter
		http.ServeFile(w, r, "view/index.html")
	})

	// Menjalankan server di port 8080
	http.ListenAndServe(":8080", nil)
}
