package middleware

import (
	"be-golang-chapter-21/impleme-http-serve/handler"
	"net/http"
)

// Middleware untuk validasi token
func TokenMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Mendapatkan token dari header Authorization
		serviceUser, db := handler.InitHandler(w)
		if db == nil {
			return
		}

		token := r.Header.Get("Token")
		if !serviceUser.TokenCheckExpire(token, w) {
			if !serviceUser.CleanExpiredTokens(db, w) {
				return
			}
			return
		}

		// Lanjutkan ke handler berikutnya jika token valid
		next.ServeHTTP(w, r)
	})
}
