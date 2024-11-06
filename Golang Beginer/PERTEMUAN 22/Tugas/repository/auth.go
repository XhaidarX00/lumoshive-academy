package repository

import (
	"main/model"
	"time"

	"github.com/google/uuid"
)

func (r *Repository) LoginRepo(user *model.Users) error {
	query := `SELECT user_id, name, username, password FROM users WHERE username=$1 AND password=$2`
	err := r.DB.QueryRow(query, user.Username, user.Password).Scan(&user.ID, &user.Name, &user.Username, &user.Password)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) RegisterRepo(user *model.Users) error {
	query := `INSERT INTO users (name, username, password) VALUES ($1, $2, $3) RETURNING id`

	var id int

	err := r.DB.QueryRow(query, user.Name, user.Username, user.Password).Scan(&id)
	if err != nil {
		return err
	}

	user.Token, err = r.GenerateTkn(id)
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) GenerateTkn(userID int) (string, error) {
	token := uuid.New().String()
	expiresAt := time.Now().Add(5 * time.Minute)

	query := "INSERT INTO tokens (user_id, token, expires_at) VALUES ($1, $2, $3)"
	_, err := r.DB.Exec(query, userID, token, expiresAt)
	if err != nil {
		return "", err
	}

	return token, nil
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
