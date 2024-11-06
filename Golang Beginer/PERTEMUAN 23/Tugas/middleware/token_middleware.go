package middlewaree

import (
	"main/handler"
	"main/service"
	"net/http"
)

// Middleware untuk validasi token
func TokenMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("Token")
		if err != nil {
			data := map[string]string{
				"ErrorMessage": "Silahkan Login Terlebih Dahulu!",
			}
			handler.RenderTemplate(w, "error.html", data)
			return
		}

		token := cookie.Value
		if err := service.ServiceF.TokenCheck(token); err != "" {
			data := map[string]string{
				"ErrorMessage": err,
			}
			handler.RenderTemplate(w, "error.html", data)
			return
		}

		// Lanjutkan ke handler berikutnya jika token valid
		next.ServeHTTP(w, r)
	})
}
