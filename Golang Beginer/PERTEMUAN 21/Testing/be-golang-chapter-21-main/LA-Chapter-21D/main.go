// // implementation form post
// package main

// import (
// 	"fmt"
// 	"log"
// 	"net/http"
// )

// func registerHandler(w http.ResponseWriter, r *http.Request) {
// 	if r.Method == http.MethodPost {
// 		// Parse form data
// 		err := r.ParseForm()
// 		if err != nil {
// 			http.Error(w, "Unable to parse form", http.StatusBadRequest)
// 			return
// 		}
// 		// Retrieve form values
// 		username := r.PostForm.Get("username")
// 		password := r.PostForm.Get("password")

// 		// Log the received data (in a real application, you should securely handle the password)
// 		fmt.Printf("Received POST request\n")
// 		fmt.Printf("Username: %s\n", username)
// 		fmt.Printf("Password: %s\n", password)

// 		// Respond to the client
// 		fmt.Fprintf(w, "Registration successful\n")
// 	} else {
// 		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
// 	}
// }

// func main() {
// 	http.HandleFunc("/register", registerHandler)

// 	// Serve the static HTML file
// 	http.Handle("/", http.FileServer(http.Dir(".")))

// 	fmt.Println("Server started at :8080")
// 	log.Fatal(http.ListenAndServe(":8080", nil))
// }

// response code
package main

import (
	"fmt"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}

func registerHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		// Parse form data
		err := r.ParseForm()
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "Bad request: unable to parse form")
			return
		}

		// Retrieve form values
		username := r.FormValue("username")
		password := r.FormValue("password")
		// Simple validation
		if username == "" || password == "" {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "Bad request: username and password cannot be empty")
			return
		}

		// Simulate successful registration
		fmt.Printf("Received registration: Username: %s, Password: %s\n", username, password)

		// Send success response
		w.WriteHeader(http.StatusCreated)
		fmt.Fprintf(w, "User registered successfully")
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "Method not allowed")
	}
}

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/register", registerHandler)

	fmt.Println("Server started at :8080")
	http.ListenAndServe(":8080", nil)
}
