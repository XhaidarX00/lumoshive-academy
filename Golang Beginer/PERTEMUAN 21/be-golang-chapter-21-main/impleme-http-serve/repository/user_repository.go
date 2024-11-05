package repository

import (
	"be-golang-chapter-21/impleme-http-serve/model"
	"database/sql"
	"fmt"
)

type UserRepository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return UserRepository{DB: db}
}

// curl -X POST http://localhost:8080/login -H "Content-Type: application/json" -d '{"username": "admin1", "password": "password123"}'
func (usr *UserRepository) Login(user *model.User) error {
	query := `SELECT username, password, role FROM users WHERE username=$1 AND password=$2`
	err := usr.DB.QueryRow(query, user.Username, user.Password).Scan(&user.Username, &user.Password, &user.Role)
	fmt.Println(err.Error())
	if err != nil {
		return err
	}
	return nil
}

// curl -X GET "http://localhost:8080/customer/customer_detail?id=1" -H "Token: 12345"
func (usr *UserRepository) UserByID(id int) (*model.User, error) {
	users := model.User{}
	query := `SELECT username, password, role FROM users WHERE id=$1`
	err := usr.DB.QueryRow(query, id).Scan(&users.Username, &users.Password)
	if err != nil {
		return nil, err
	}
	return &users, nil
}
