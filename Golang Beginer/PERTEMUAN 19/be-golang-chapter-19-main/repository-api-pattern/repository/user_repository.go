package repository

import (
	"be-golang-chapter-19/repository-api-pattern/model"
	"database/sql"
)

type CustomerRepositoryDB struct {
	DB *sql.DB
}

func NewCustomerRepository(db *sql.DB) CustomerRepositoryDB {
	return CustomerRepositoryDB{DB: db}
}

func (r *CustomerRepositoryDB) Create(customer *model.Customer) error {
	query := `INSERT INTO customers (username, password, email) VALUES ($1, $2, $3) RETURNING id`
	err := r.DB.QueryRow(query, customer.Username, customer.Password, customer.Email).Scan(&customer.ID)
	if err != nil {
		return err
	}

	return nil
}

func (r *CustomerRepositoryDB) GetAll() (*[]model.Customer, error) {
	query := `SELECT id, username, password, email FROM customers`
	rows, err := r.DB.Query(query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	customers := []model.Customer{}

	for rows.Next() {
		var customer model.Customer
		rows.Scan(&customer.ID, &customer.Username, &customer.Password, &customer.Email)

		customers = append(customers, customer)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &customers, nil
}

func (r *CustomerRepositoryDB) GetCustumerLogin(customer model.Customer) (*model.Customer, error) {
	query := `SELECT id, username, password, email FROM customers WHERE username=$1 AND password=$2`
	var customerResponse model.Customer
	err := r.DB.QueryRow(query, customer.Username, customer.Password).Scan(&customerResponse.ID, &customerResponse.Username, &customerResponse.Password, &customerResponse.Email)

	if err != nil {
		return nil, err
	}

	return &customerResponse, nil
}
