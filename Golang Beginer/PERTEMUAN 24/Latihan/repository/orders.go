package repository

import (
	"fmt"
	"latihan/model/orders"
)

func (r *Repository) GetOrderDataRepo(data *[]orders.Order) error {
	rows, err := r.DB.Query("SELECT id, customer_id, payment_methode, total_amount, discount, final_amount, order_date, status FROM orders ORDER BY id")
	if err != nil {
		fmt.Println(err.Error())
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
			panic(err)
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
