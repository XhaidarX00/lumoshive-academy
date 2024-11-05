//

// server file
package main

import (
	"net/http"
)

func main() {
	// fs := http.FileServer(http.Dir("./view"))

	// http.Handle("/", fs)

	// if err := http.ListenAndServe(":8080", nil); err != nil {
	// 	panic(err)
	// }

	// Mengatur route untuk melayani file statis
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		// Melayani file HTML ke ResponseWriter
		http.ServeFile(w, r, "index.html")
	})

	// Menjalankan server di port 8080
	http.ListenAndServe(":8080", nil)

}
