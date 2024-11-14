package repository

import (
	"latihan/model"
)

func (t *Repo) AddTransactionRepo(data *model.AddTransaction) error {
	query := `
	INSERT INTO transaction (name, email, phone_number, message, event_id, status_order)
	VALUES
	($1, $2, $3, $4, $5, $6) RETURNING id
	`
	err := t.DB.QueryRow(query, data.Name, data.Email, data.Phone, data.Message, data.Event_id, data.Status_order).Scan(&data.ID)
	if err != nil {
		return err
	}

	// query2 := `SELECT id, name, email, phone_number, message, event_id, status_order FROM transaction WHERE id = $1`
	// err = t.DB.QueryRow(query2, data.ID).Scan(&data.ID, &data.Name, &data.Email, &data.Phone, &data.Message, &data.Event_id, &data.Status_order)
	// if err != nil {
	// 	return err
	// }

	return nil
}
