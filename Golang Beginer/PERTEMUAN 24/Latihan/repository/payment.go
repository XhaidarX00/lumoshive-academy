// repository/payment_repository.go
package repository

import (
	"database/sql"
	"errors"
	"latihan/model/payment"
)

func (r *Repository) Create(payment *payment.Payment) error {
	query := `INSERT INTO payments (name, photo, is_active, created_at, updated_at) 
              VALUES ($1, $2, $3, NOW(), NOW()) RETURNING id`
	err := r.DB.QueryRow(query, payment.Name, payment.Photo, payment.IsActive).Scan(&payment.ID)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) GetAll() ([]payment.Payment, error) {
	query := `SELECT id, name, photo, is_active, created_at, updated_at, deleted_at FROM payments WHERE deleted_at IS NULL`
	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var payments []payment.Payment
	for rows.Next() {
		var payment payment.Payment
		if err := rows.Scan(&payment.ID, &payment.Name, &payment.Photo, &payment.IsActive, &payment.CreatedAt, &payment.UpdatedAt, &payment.DeletedAt); err != nil {
			return nil, err
		}
		payments = append(payments, payment)
	}
	return payments, nil
}

func (r *Repository) GetByID(id int) (*payment.Payment, error) {
	query := `SELECT id, name, photo, is_active, created_at, updated_at, deleted_at FROM payments WHERE id=$1 AND deleted_at IS NULL`
	var payment payment.Payment
	err := r.DB.QueryRow(query, id).Scan(&payment.ID, &payment.Name, &payment.Photo, &payment.IsActive, &payment.CreatedAt, &payment.UpdatedAt, &payment.DeletedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("payment not found")
		}
		return nil, err
	}
	return &payment, nil
}

func (r *Repository) Update(payment *payment.Payment) error {
	query := `UPDATE payments SET name=$1, photo=$2, is_active=$3, updated_at=NOW() WHERE id=$4 AND deleted_at IS NULL`
	_, err := r.DB.Exec(query, payment.Name, payment.Photo, payment.IsActive, payment.ID)
	return err
}

func (r *Repository) Delete(id int) error {
	query := `UPDATE payments SET deleted_at=NOW() WHERE id=$1`
	_, err := r.DB.Exec(query, id)
	return err
}
