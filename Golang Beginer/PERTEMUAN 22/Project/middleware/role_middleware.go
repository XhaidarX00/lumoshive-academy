package middleware

import (
	"encoding/json"
	"main/handler"
	"main/model"
	"net/http"
)

// Middleware untuk validasi token
func RoleMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("token")
		db, servicef := handler.InitDBHandler(w)
		if db == nil {
			return
		}

		role, err := servicef.GetRoleService(token)
		if err != nil {
			badResponse := model.ResponseError{
				StatusCode: http.StatusBadRequest,
				Message:    err.Error(),
			}
			json.NewEncoder(w).Encode(badResponse)
			return
		}

		if !servicef.RoleCheckAccService(role, w, r) {
			return
		}

		// Lanjutkan ke handler berikutnya jika token valid
		next.ServeHTTP(w, r)
	})
}
