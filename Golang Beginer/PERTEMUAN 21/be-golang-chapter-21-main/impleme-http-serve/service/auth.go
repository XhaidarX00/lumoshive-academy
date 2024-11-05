package service

import (
	"database/sql"
	"net/http"
)

// fungsi untuk check token
func (usr *UserService) TokenCheckExpire(token string, w http.ResponseWriter) bool {
	return usr.RepoUser.TokenCheckExp(token, w)
}

// fungsi untuk membersihkan token yang sudah kadaluarsa
func (usr *UserService) CleanExpiredTokens(db *sql.DB, w http.ResponseWriter) bool {
	return usr.RepoUser.CleanExpiredTkn(db, w)
}

// fungsi untuk memeriksa role
func (usr *UserService) RoleCheckAccses(role string, w http.ResponseWriter, r *http.Request) bool {
	return usr.RepoUser.RoleCheckAcc(role, w, r)
}

// fungsi untuk generate token
func (usr *UserService) GenerateToken(db *sql.DB, userID int, w http.ResponseWriter) string {
	token := usr.RepoUser.GenerateTkn(db, userID, w)
	if token == "" {
		return ""
	}
	return token
}
