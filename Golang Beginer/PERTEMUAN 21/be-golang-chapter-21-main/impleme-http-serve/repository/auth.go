package repository

import (
	"be-golang-chapter-21/impleme-http-serve/model"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
)

// Fungsi untuk membersihkan token yang sudah kadaluarsa
func (usr *UserRepository) CleanExpiredTkn(db *sql.DB, w http.ResponseWriter) bool {
	query := "DELETE FROM tokens WHERE expires_at < $1"
	_, err := db.Exec(query, time.Now())
	if err != nil {
		// Jika gagal, kirim respons error dengan format yang diminta
		badResponse := model.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    "Failed to clean expired tokens",
			Data:       nil,
		}
		json.NewEncoder(w).Encode(badResponse)
		log.Printf("Failed to clean expired tokens: %v", err)
		return false
	}

	log.Println("Expired tokens cleaned successfully")
	successResponse := model.Response{
		StatusCode: http.StatusOK,
		Message:    "Expired tokens cleaned successfully",
		Data:       nil,
	}
	json.NewEncoder(w).Encode(successResponse)

	return true
}

func (usr *UserRepository) TokenCheckExp(token string, w http.ResponseWriter) bool {
	if token == "" {
		badResponse := model.Response{
			StatusCode: http.StatusUnauthorized,
			Message:    "Token required",
			Data:       nil,
		}
		json.NewEncoder(w).Encode(badResponse)
		return false
	}

	// Memeriksa token di database
	var expiresAt time.Time
	query := "SELECT expires_at FROM tokens WHERE token = $1"
	err := usr.DB.QueryRow(query, token).Scan(&expiresAt)
	if err != nil {
		badResponse := model.Response{
			StatusCode: http.StatusUnauthorized,
			Message:    "Invalid or expired token",
			Data:       nil,
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

func (usr *UserRepository) RoleCheckAcc(role string, w http.ResponseWriter, r *http.Request) bool {
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

func (usr *UserRepository) GenerateTkn(db *sql.DB, userID int, w http.ResponseWriter) string {
	token := uuid.New().String()

	expiresAt := time.Now().Add(5 * time.Minute)

	query := "INSERT INTO tokens (user_id, token, expires_at) VALUES ($1, $2, $3)"
	_, err := db.Exec(query, userID, token, expiresAt)
	if err != nil {
		badResponse := model.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    "Failed to generate token",
			Data:       nil,
		}

		json.NewEncoder(w).Encode(badResponse)
		return ""
	}

	return token
}
