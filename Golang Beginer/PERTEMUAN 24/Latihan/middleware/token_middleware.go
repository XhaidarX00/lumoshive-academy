package middlewaree

import (
	pagehandler "latihan/controller/pageHandler"
	"net/http"
)

// Middleware untuk validasi token
func (m *Middlewaree) TokenMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("Token")
		// token := r.Header.Get("Token")
		if err != nil {
			// if token == "" {
			data := map[string]string{
				"ErrorMessage": "Silahkan Login Terlebih Dahulu!",
			}
			pagehandler.RenderTemplate(w, "error.html", data)
			return
		}

		token := cookie.Value
		if err := m.svc.TokenCheck(token); err != "" {
			data := map[string]string{
				"ErrorMessage": err,
			}
			pagehandler.RenderTemplate(w, "error.html", data)
			return
		}

		// Lanjutkan ke handler berikutnya jika token valid
		next.ServeHTTP(w, r)
	})
}
