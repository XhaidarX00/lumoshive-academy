package repository

import (
	"encoding/json"
	"main/model"
	"net/http"
	"time"

	"github.com/google/uuid"
)

func (r *Repository) LoginRepo(user *model.User) error {
	query := `SELECT user_id, username, password, email, role FROM users WHERE email=$1 AND password=$2`
	err := r.DB.QueryRow(query, user.Email, user.Password).Scan(&user.ID, &user.Username, &user.Password, &user.Email, &user.Role)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) RegisterRepo(user *model.User) error {
	query := `INSERT INTO users (username, password, email, role) VALUES ($1, $2, $3, $4)`
	_, err := r.DB.Exec(query, user.Username, user.Password, user.Email, user.Role)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) GenerateTkn(userID int, w http.ResponseWriter) string {
	token := uuid.New().String()

	expiresAt := time.Now().Add(5 * time.Minute)

	query := "INSERT INTO tokens (user_id, token, expires_at) VALUES ($1, $2, $3)"
	_, err := r.DB.Exec(query, userID, token, expiresAt)
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

func (r *Repository) GetRoleRepo(token string) (string, error) {
	var role string
	query := `SELECT u.role FROM tokens t JOIN users u ON u.user_id = t.user_id WHERE t.token = $1;`
	err := r.DB.QueryRow(query, token).Scan(&role)
	if err != nil {
		return "", err
	}
	return role, nil
}
