package repository

import (
	"latihan/library"
	"latihan/model/orders"

	"go.uber.org/zap"
)

func (r *Repository) GetOrderDataRepo(data *[]orders.Order) error {
	rows, err := r.DB.Query("SELECT id, customer_id, payment_methode, total_amount, discount, final_amount, order_date, status FROM orders ORDER BY id")
	if err != nil {
		r.Logger.Error("Error GetOrderDataRepo", zap.Error(err))
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var order orders.Order
		if err := rows.Scan(
			&order.ID,
			&order.Customer_id,
			&order.Payment_methode,
			&order.Total_amount,
			&order.Discount,
			&order.Final_amount,
			&order.OrderDate,
			&order.Status,
		); err != nil {
			r.Logger.Error("Error GetOrderDataRepo", zap.Error(err))
			return err
		}

		order.OrderDate, err = library.ChangeFormatTime(order.OrderDate)
		if err != nil {
			r.Logger.Error("Error GetOrderDataRepo", zap.Error(err))
			return err
		}

		name, err := r.GetCustomerByIDRepo(order.Customer_id)
		if err != nil {
			return err
		}

		order.Name_customer = name
		*data = append(*data, order)
	}

	return nil
}

func (r *Repository) GetOrderDetailRepo(order *orders.Order) error {
	err := r.DB.QueryRow(`SELECT 
	id, 
	customer_id, 
	payment_methode, 
	total_amount, 
	discount, 
	final_amount, 
	order_date, 
	status 
	FROM orders 
	WHERE id = $1
	ORDER BY id`, order.ID).Scan(
		&order.ID,
		&order.Customer_id,
		&order.Payment_methode,
		&order.Total_amount,
		&order.Discount,
		&order.Final_amount,
		&order.OrderDate,
		&order.Status,
	)

	if err != nil {
		r.Logger.Error("Error GetOrderDetail : ", zap.Error(err))
		return err
	}

	return nil
}
