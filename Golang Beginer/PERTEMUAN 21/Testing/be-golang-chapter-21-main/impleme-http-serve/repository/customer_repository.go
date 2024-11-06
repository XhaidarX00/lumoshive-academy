package repository

import (
	"be-golang-chapter-21/impleme-http-serve/model"
	"database/sql"
)

type CustomerRepository struct {
	DB *sql.DB
}

func NewCustomerRepository(db *sql.DB) CustomerRepository {
	return CustomerRepository{DB: db}
}

func (cr *CustomerRepository) Login(customer *model.Customer) error {
	query := `SELECT username, password, email FROM customers WHERE username=$1 AND password=$2`
	err := cr.DB.QueryRow(query, customer.Username, customer.Password).Scan(&customer.Username, &customer.Password, &customer.Email)
	if err != nil {
		return err
	}
	return nil
}

func (cr *CustomerRepository) CustomerByID(id int) (*model.Customer, error) {
	customer := model.Customer{}
	query := `SELECT username, password, email FROM customers WHERE id=$1`
	err := cr.DB.QueryRow(query, id).Scan(&customer.Username, &customer.Password, &customer.Email)
	if err != nil {
		return nil, err
	}
	return &customer, nil
}
