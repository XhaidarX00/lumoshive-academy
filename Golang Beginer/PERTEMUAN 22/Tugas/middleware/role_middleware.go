package middleware

import (
	"main/library"
	"main/service"
	"net/http"
)

// Middleware untuk validasi token
func RoleMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("token")
		role, err := service.ServiceF.GetRoleService(token)
		if err != nil {
			library.ResponseToJson(w, err.Error(), nil)
			return
		}

		if !service.ServiceF.RoleCheckAcc(role, w, r) {
			return
		}

		// Lanjutkan ke handler berikutnya jika token valid
		next.ServeHTTP(w, r)
	})
}
