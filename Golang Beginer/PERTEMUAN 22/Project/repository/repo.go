package repository

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"main/model"
	"main/utils"
	"net/http"
	"time"
)

type Repository struct {
	DB *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return Repository{DB: db}
}

func (r *Repository) TokenCheckRepo(db *sql.DB, token string, w http.ResponseWriter) bool {
	if token == "" {
		badResponse := model.ResponseError{
			StatusCode: http.StatusUnauthorized,
			Message:    "Token required",
		}
		json.NewEncoder(w).Encode(badResponse)
		return false
	}

	var expiresAt time.Time
	query := "SELECT expires_at FROM tokens WHERE token = $1"
	err := db.QueryRow(query, token).Scan(&expiresAt)
	if err != nil {
		badResponse := model.ResponseError{
			StatusCode: http.StatusUnauthorized,
			Message:    "Invalid or expired token",
		}
		json.NewEncoder(w).Encode(badResponse)
		return false
	}

	// Memeriksa apakah token sudah kadaluarsa
	if time.Now().After(expiresAt) {
		badResponse := model.Response{
			StatusCode: http.StatusUnauthorized,
			Message:    "Token has expired",
			Data:       nil,
		}
		json.NewEncoder(w).Encode(badResponse)
		return false
	}

	return true
}

// Fungsi untuk membersihkan token yang sudah kadaluarsa
func (r *Repository) CleanExpiredTokensRepo(db *sql.DB, w http.ResponseWriter) bool {
	query := "DELETE FROM tokens WHERE expires_at < $1"
	_, err := db.Exec(query, time.Now())
	if err != nil {
		badResponse := model.ResponseError{
			StatusCode: http.StatusUnauthorized,
			Message:    fmt.Sprintf("Failed to clean expired tokens: %v", err),
		}
		json.NewEncoder(w).Encode(badResponse)
		return false
	}

	badResponse := model.Response{
		StatusCode: http.StatusUnauthorized,
		Message:    "Expired tokens cleaned successfully",
		Data:       nil,
	}
	json.NewEncoder(w).Encode(badResponse)
	return true
}

func (r *Repository) CleanExpiredTokensRepo2(db *sql.DB) bool {
	query := "DELETE FROM tokens WHERE expires_at < $1"
	_, err := db.Exec(query, time.Now())
	if err != nil {
		utils.ErrorMessage(err.Error())
		return false
	}
	return true
}

func (rp *Repository) RoleCheckAccRepo(role string, w http.ResponseWriter, r *http.Request) bool {
	// Check role permissions based on HTTP method
	switch role {
	case "dev":
		if r.Method != http.MethodPut {
			badResponse := model.Response{
				StatusCode: http.StatusForbidden,
				Message:    "Forbidden: Only 'PUT' method is allowed for dev role",
				Data:       nil,
			}
			w.WriteHeader(http.StatusForbidden)
			json.NewEncoder(w).Encode(badResponse)
			return false
		}

		return true
	case "admin":
		return true
	default:
		badResponse := model.Response{
			StatusCode: http.StatusForbidden,
			Message:    "Forbidden: Unrecognized role",
			Data:       nil,
		}
		w.WriteHeader(http.StatusForbidden)
		json.NewEncoder(w).Encode(badResponse)
		return false
	}
}
