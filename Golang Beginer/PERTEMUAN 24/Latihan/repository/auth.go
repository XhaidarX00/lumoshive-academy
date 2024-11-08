package repository

import (
	"latihan/model/customers"
	"time"

	"github.com/google/uuid"
)

func (r *Repository) LoginRepo(user *customers.Customer) error {
	query := `SELECT u.id, u.name, u.username, u.password, t.token 
	FROM customers u
	JOIN tokens t ON t.customer_id = u.id
	WHERE u.username = $1 AND u.password = $2`
	err := r.DB.QueryRow(query, user.Username, user.Password).Scan(&user.ID, &user.Name, &user.Username, &user.Password, &user.Token)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) RegisterRepo(user *customers.Customer) error {
	query := `INSERT INTO customers (name, username, password, phone_Number) VALUES ($1, $2, $3, $4) RETURNING id`

	var id int

	err := r.DB.QueryRow(query, user.Name, user.Username, user.Password, user.Phone_Number).Scan(&id)
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

	query := "INSERT INTO tokens (customer_id, token, expires_at) VALUES ($1, $2, $3)"
	_, err := r.DB.Exec(query, userID, token, expiresAt)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (r *Repository) GetRoleRepo(token string) (string, error) {
	var role string
	query := `SELECT u.role FROM tokens t JOIN costumers u ON u.user_id = t.user_id WHERE t.token = $1;`
	err := r.DB.QueryRow(query, token).Scan(&role)
	if err != nil {
		return "", err
	}
	return role, nil
}

func (r *Repository) GetCustomerByIDRepo(id int) (string, error) {
	var name string
	query := `SELECT u.name
	FROM customers u
	JOIN orders o ON o.customer_id = u.id 
	WHERE o.customer_id = $1
	LIMIT 1
`
	err := r.DB.QueryRow(query, id).Scan(&name)
	if err != nil {
		return "", err
	}
	return name, nil
}
