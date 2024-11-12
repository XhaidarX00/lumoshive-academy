// repository/payment_repository.go
package repository

import (
	"database/sql"
	"errors"
	"latihan/model/payment"

	"go.uber.org/zap"
)

func (r *Repository) Create(payment *payment.Payment) error {
	query := `INSERT INTO payments (name, photo, is_active, created_at, updated_at) 
              VALUES ($1, $2, $3, NOW(), NOW()) RETURNING id`
	err := r.DB.QueryRow(query, payment.Name, payment.Photo, payment.IsActive).Scan(&payment.ID)
	if err != nil {
		r.Logger.Error("Error CreatePayment", zap.Error(err))
		return err
	}
	return nil
}

func (r *Repository) GetAll() ([]payment.Payment, error) {
	query := `SELECT id, name, photo, is_active, created_at, updated_at, deleted_at FROM payments WHERE deleted_at IS NULL`
	rows, err := r.DB.Query(query)
	if err != nil {
		r.Logger.Error("Error GetAllRepo", zap.Error(err))
		return nil, err
	}
	defer rows.Close()

	var payments []payment.Payment
	for rows.Next() {
		var payment payment.Payment
		if err := rows.Scan(&payment.ID, &payment.Name, &payment.Photo, &payment.IsActive, &payment.CreatedAt, &payment.UpdatedAt, &payment.DeletedAt); err != nil {
			r.Logger.Error("Error GetAllPayment", zap.Error(err))
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
		r.Logger.Error("Error GetByID", zap.Error(err))
		if err == sql.ErrNoRows {
			r.Logger.Error("Terjadi Kesalahan Repository Database",
				zap.String("Name Repository", "Repository.Payment"),
				zap.String("Name Func", "GetByID"),
				zap.String("Query", query),
				zap.Error(err),
			)
			return nil, errors.New("payment not found")
		}
		return nil, err
	}
	return &payment, nil
}

func (r *Repository) Update(payment *payment.Payment) error {
	query := `UPDATE payments SET name=$1, photo=$2, is_active=$3, updated_at=NOW() WHERE id=$4 AND deleted_at IS NULL`
	_, err := r.DB.Exec(query, payment.Name, payment.Photo, payment.IsActive, payment.ID)
	if err != nil {
		r.Logger.Error("Error UpdatePayment", zap.Error(err))
	}
	return err
}

func (r *Repository) Delete(id int) error {
	query := `UPDATE payments SET deleted_at=NOW() WHERE id=$1`
	_, err := r.DB.Exec(query, id)
	if err != nil {
		r.Logger.Error("Error DeletePayment", zap.Error(err))
	}
	return err
}
