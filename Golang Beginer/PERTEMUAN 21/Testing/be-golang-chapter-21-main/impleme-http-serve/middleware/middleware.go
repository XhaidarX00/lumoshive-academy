package middleware

import (
	"be-golang-chapter-21/impleme-http-serve/model"
	"encoding/json"
	"net/http"
)

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("token")
		if authHeader != "12345" {
			badResponse := model.Response{
				StatusCode: http.StatusUnauthorized,
				Message:    "Unauthorized",
				Data:       nil,
			}
			json.NewEncoder(w).Encode(badResponse)
			return
		}

		// Melanjutkan ke handler berikutnya
		next.ServeHTTP(w, r)
	})
}
