package main

import (
	"fmt"
	"net/http"
)

func setCookieHandler(w http.ResponseWriter, r *http.Request) {
	cookie := http.Cookie{
		Name:  "username",
		Value: "golang_user",
		Path:  "/",
	}
	http.SetCookie(w, &cookie)
	fmt.Fprintln(w, "Cookie telah disetel")
}

func getCookieHandler(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("username")
	if err != nil {
		fmt.Fprintln(w, "Gagal mengambil cookie:", err)
		return
	}
	fmt.Fprintln(w, "Nilai cookie:", cookie.Value)
}

func main() {
	http.HandleFunc("/setcookie", setCookieHandler)
	http.HandleFunc("/getcookie", getCookieHandler)
	http.ListenAndServe(":8080", nil)
}
