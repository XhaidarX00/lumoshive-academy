package main

import (
	"fmt"
	"net/http"
	"time"
)

// LoggingMiddleware adalah middleware untuk logging
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// Proses permintaan
		next.ServeHTTP(w, r)

		// Logging setelah permintaan selesai
		fmt.Printf("[%s] %s %s %v\n", r.Method, r.URL.Path, r.RemoteAddr, time.Since(start))
	})
}

// AuthMiddleware adalah middleware untuk autentikasi sederhana
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Cek keberadaan token atau sesi yang valid
		token := r.Header.Get("Authorization")
		if token != "token_valid" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// Lanjutkan ke handler berikutnya jika autentikasi berhasil
		next.ServeHTTP(w, r)
	})
}

// Handler untuk halaman utama
func mainHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Halaman Utama")
}

// Handler untuk sumber daya yang dilindungi
func protectedHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Sumber Daya Dilindungi")
}

func main() {
	// Inisialisasi router default
	mux := http.NewServeMux()

	// Tambahkan middleware logging untuk semua route
	mux.Handle("/", LoggingMiddleware(http.HandlerFunc(mainHandler)))
	mux.Handle("/protected", LoggingMiddleware(AuthMiddleware(http.HandlerFunc(protectedHandler))))

	// Mulai server dengan router yang telah dikonfigurasi
	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	fmt.Println("Server berjalan di http://localhost:8080")
	server.ListenAndServe()
}
