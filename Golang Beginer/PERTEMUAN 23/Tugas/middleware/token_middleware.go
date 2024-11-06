package middlewaree

import (
	"main/service"
	"net/http"
)

// Middleware untuk validasi token
func TokenMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.URL.Query().Get("token")
		if !service.ServiceF.TokenCheck(token, w) {
			return
		}

		// Lanjutkan ke handler berikutnya jika token valid
		next.ServeHTTP(w, r)
	})
}
