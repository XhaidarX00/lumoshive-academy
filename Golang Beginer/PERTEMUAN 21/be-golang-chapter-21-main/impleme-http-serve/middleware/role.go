package middleware

import (
	"be-golang-chapter-21/impleme-http-serve/handler"
	"net/http"
)

// Role function to validate role permissions
func Role(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		serviceUser, db := handler.InitHandler(w)
		if db == nil {
			return
		}

		role := r.Header.Get("Role")
		if !serviceUser.RoleCheckAccses(role, w, r) {
			return
		}

		next.ServeHTTP(w, r)
	})
}
