package middleware

import (
	"main/handler"
	"net/http"
)

// Middleware untuk validasi token
func TokenMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("token")
		db, servicef := handler.InitDBHandler(w)
		if db == nil {
			return
		}

		if !servicef.TokenCheckService(db, token, w) {
			return
		}

		// Lanjutkan ke handler berikutnya jika token valid
		next.ServeHTTP(w, r)
	})
}
