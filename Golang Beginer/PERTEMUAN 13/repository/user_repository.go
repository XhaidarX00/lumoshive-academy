package repository

import (
	"database/sql"
	"log"
	utils "main/Utils"
	"main/model"
	"time"
)

type CustomerRepositoryDB struct {
	DB *sql.DB
}

func NewCustomRepository(db *sql.DB) CustomerRepository {
	return &CustomerRepositoryDB{DB: db}
}

// -- Created: Customers
func (r *CustomerRepositoryDB) CreateCustomers(customer *model.Customers) (uint16, error) {
	query := "INSERT INTO Customers (name, email, phone_number, created_at) VALUES ($1, $2, $3, $4) RETURNING customer_id"
	err := r.DB.QueryRow(query, customer.Name, customer.Email, customer.Phone_number, customer.Created_at).Scan(&customer.Customer_id)
	if err != nil {
		return 0, err
	}

	return customer.Customer_id, nil
}

// -- Created: Drivers
func (r *CustomerRepositoryDB) CreateDrivers(driver *model.Drivers) (uint16, error) {
	query := "INSERT INTO Drivers (name, phone_number, vehicle_type, created_at) VALUES ($1, $2, $3, $4) RETURNING driver_id"
	err := r.DB.QueryRow(query, driver.Name, driver.Phone_number, driver.Vehicle_type, driver.Created_at).Scan(&driver.Driver_id)
	if err != nil {
		return 0, err
	}

	return driver.Driver_id, nil
}

// -- Created: Orders
func (r *CustomerRepositoryDB) CreateOrders(order *model.Orders) (uint16, error) {
	// Memulai transaksi
	tx, err := r.DB.Begin()
	if err != nil {
		return 0, err
	}

	orderQuery := "INSERT INTO Orders (customer_id, driver_id, order_date, pickup_location, dropoff_location, total_fare) VALUES ($1, $2, $3, $4, $5, $6) RETURNING order_id"
	err = tx.QueryRow(orderQuery, order.Customer_id, order.Driver_id, order.Order_date, order.Pickup_location, order.Dropoff_location, order.Total_fare).Scan(&order.Order_id)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	statusQuery := "INSERT INTO OrderStatus (order_id, status, updated_at) VALUES ($1, $2, $3)"
	_, err = tx.Exec(statusQuery, order.Order_id, "ongoing", time.Now())
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	// Menyelesaikan transaksi
	err = tx.Commit()
	if err != nil {
		return 0, err
	}

	return order.Order_id, nil
}

// func (r *CustomerRepositoryDB) CreateOrders(order *model.Orders) (uint16, error) {
// 	query := "INSERT INTO customer (customer_id, driver_id, order_date, pickup_location, dropoff_location, total_order) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id"
// 	err := r.DB.QueryRow(query, order.Customer_id, order.Driver_id, order.Order_date, order.Pickup_location, order.Dropoff_location, order.Total_order).Scan(&order.ID)
// 	if err != nil {
// 		return 0, err
// 	}

// 	return order.ID, err
// }

// -- Created: OrderStatus

// -- View: Total order setiap bulan
// "SELECT DATE_TRUNC('month', order_date) AS month, COUNT(order_id) AS total_orders FROM Orders GROUP BY month;"
func (r *CustomerRepositoryDB) ViewMonthlyOrders() (interface{}, error) {
	query := "SELECT DATE_TRUNC('month', order_date) AS month, COUNT(order_id) AS total_orders FROM Orders GROUP BY month"

	rows, err := r.DB.Query(query)
	if err != nil {
		return "", nil
	}

	var MonthlyOrders []struct {
		Month        string
		Total_orders int
	}

	for rows.Next() {
		MonthlyOrder := struct {
			Month        string
			Total_orders int
		}{}

		err := rows.Scan(&MonthlyOrder.Month, &MonthlyOrder.Total_orders)
		if err != nil {
			log.Fatal(err)
		}
		MonthlyOrders = append(MonthlyOrders, MonthlyOrder)
	}
	data := utils.ConvertSliceToMap(MonthlyOrders)
	keys := utils.GetStructKeys(MonthlyOrders[0], []string{})
	utils.DisplayData("Data Top Driver Tiap Bulan", data, keys)
	return MonthlyOrders, nil
}

// -- View: Customer yang sering order tiap bulan
// "SELECT DATE_TRUNC('month', o.order_date) AS month, o.customer_id, c.name, COUNT(o.order_id) AS total_orders FROM Orders o JOIN customers c ON o.customer_id = c.customer_id GROUP BY month, o.customer_id, c.name ORDER BY month, total_orders DESC"
func (r *CustomerRepositoryDB) ViewMonthlyCustomerOrders() (interface{}, error) {
	query := "SELECT DATE_TRUNC('month', o.order_date) AS month, o.customer_id, c.name, COUNT(o.order_id) AS total_orders FROM Orders o JOIN customers c ON o.customer_id = c.customer_id GROUP BY month, o.customer_id, c.name ORDER BY month, total_orders DESC"

	rows, err := r.DB.Query(query)
	if err != nil {
		return "", nil
	}

	var MonthlyCustomerOrders []struct {
		Month        string
		Customer_id  int
		Name         string
		Total_orders int
	}

	for rows.Next() {
		MonthlyOrder := struct {
			Month        string
			Customer_id  int
			Name         string
			Total_orders int
		}{}

		err := rows.Scan(&MonthlyOrder.Month, &MonthlyOrder.Customer_id, &MonthlyOrder.Name, &MonthlyOrder.Total_orders)
		if err != nil {
			log.Fatal(err)
		}
		MonthlyCustomerOrders = append(MonthlyCustomerOrders, MonthlyOrder)
	}
	data := utils.ConvertSliceToMap(MonthlyCustomerOrders)
	keys := utils.GetStructKeys(MonthlyCustomerOrders[0], []string{})
	utils.DisplayData("Data Top Driver Tiap Bulan", data, keys)
	return MonthlyCustomerOrders, nil
}

// -- View: Daerah dengan banyak order
// "SELECT L.name AS pickup_location, COUNT(O.order_id) AS total_orders FROM Orders O JOIN Locations L ON O.pickup_location = L.location_id GROUP BY L.name ORDER BY total_orders DESC"
func (r *CustomerRepositoryDB) GetMonthlyCustomerOrders() (interface{}, error) {
	query := "SELECT L.name AS pickup_location, COUNT(O.order_id) AS total_orders FROM Orders O JOIN Locations L ON O.pickup_location = L.location_id GROUP BY L.name ORDER BY total_orders DESC"

	rows, err := r.DB.Query(query)
	if err != nil {
		return "", err
	}

	var GetMonthlyCustomerOrders []struct {
		Pickup_location string
		Total_orders    int
	}

	for rows.Next() {
		MonthlyOrder := struct {
			Pickup_location string
			Total_orders    int
		}{}

		err := rows.Scan(&MonthlyOrder.Pickup_location, &MonthlyOrder.Total_orders)
		if err != nil {
			log.Fatal(err)
		}
		GetMonthlyCustomerOrders = append(GetMonthlyCustomerOrders, MonthlyOrder)
	}
	data := utils.ConvertSliceToMap(GetMonthlyCustomerOrders)
	keys := utils.GetStructKeys(GetMonthlyCustomerOrders[0], []string{})
	utils.DisplayData("Data Top Driver Tiap Bulan", data, keys)
	return GetMonthlyCustomerOrders, nil
}

// -- View: Pukul berapa saja order yang ramai dan sepi
// "SELECT TO_CHAR(order_date, 'HH24:MI') AS hour_minute, COUNT(order_id) AS total_orders FROM Orders GROUP BY hour_minute ORDER BY total_orders DESC"
func (r *CustomerRepositoryDB) ViewOrderPeakTimes() (interface{}, error) {
	query := "SELECT TO_CHAR(order_date, 'HH24:MI') AS hour_minute, COUNT(order_id) AS total_orders FROM Orders GROUP BY hour_minute ORDER BY total_orders DESC"

	rows, err := r.DB.Query(query)
	if err != nil {
		return "", err
	}

	var ViewOrderPeakTimes []struct {
		Hour_minute  string
		Total_orders int
	}

	for rows.Next() {
		MonthlyOrder := struct {
			Hour_minute  string
			Total_orders int
		}{}

		err := rows.Scan(&MonthlyOrder.Hour_minute, &MonthlyOrder.Total_orders)
		if err != nil {
			log.Fatal(err)
		}
		ViewOrderPeakTimes = append(ViewOrderPeakTimes, MonthlyOrder)
	}
	data := utils.ConvertSliceToMap(ViewOrderPeakTimes)
	keys := utils.GetStructKeys(ViewOrderPeakTimes[0], []string{})
	utils.DisplayData("Data Top Driver Tiap Bulan", data, keys)
	return ViewOrderPeakTimes, nil
}

// -- View: Jumlah customer yang masih login dan belum logout
// "SELECT cl.customer_id, c.name,c.phone_number,c.email, COUNT(CASE WHEN cl.logout_time IS NULL THEN 1 END) AS total_logins,COUNT(CASE WHEN cl.logout_time IS NOT NULL THEN 1 END) AS total_logouts FROM CustomerLogins cl JOIN Customers c ON c.customer_id = cl.customer_id GROUP BY cl.customer_id, c.name, c.phone_number, c.email"
func (r *CustomerRepositoryDB) ViewCustomersSession() (interface{}, error) {
	query := "SELECT cl.customer_id, c.name,c.phone_number,c.email, COUNT(CASE WHEN cl.logout_time IS NULL THEN 1 END) AS total_logins,COUNT(CASE WHEN cl.logout_time IS NOT NULL THEN 1 END) AS total_logouts FROM CustomerLogins cl JOIN Customers c ON c.customer_id = cl.customer_id GROUP BY cl.customer_id, c.name, c.phone_number, c.email"

	rows, err := r.DB.Query(query)
	if err != nil {
		return "", err
	}

	var ViewCustomersSession []struct {
		Customer_id   int
		Name          string
		Phone_number  string
		Email         string
		Total_logins  string
		Total_logouts string
	}

	for rows.Next() {
		MonthlyOrder := struct {
			Customer_id   int
			Name          string
			Phone_number  string
			Email         string
			Total_logins  string
			Total_logouts string
		}{}

		err := rows.Scan(&MonthlyOrder.Customer_id, &MonthlyOrder.Name, &MonthlyOrder.Phone_number, &MonthlyOrder.Email, &MonthlyOrder.Total_logins, &MonthlyOrder.Total_logouts)
		if err != nil {
			log.Fatal(err)
		}
		ViewCustomersSession = append(ViewCustomersSession, MonthlyOrder)
	}
	data := utils.ConvertSliceToMap(ViewCustomersSession)
	keys := utils.GetStructKeys(ViewCustomersSession[0], []string{})
	utils.DisplayData("Data Top Driver Tiap Bulan", data, keys)
	return ViewCustomersSession, nil
}

// -- View: Driver yang rajin mengambil order setiap bulan
// "SELECT DATE_TRUNC('month', o.order_date) AS month, d.driver_id, d.name, COUNT(CASE WHEN os.status = 'completed' THEN o.order_id END) AS total_success_orders, COUNT(CASE WHEN os.status = 'cancelled' THEN o.order_id END) AS total_cancelled_orders,COUNT(o.order_id) AS total_orders FROM Orders o JOIN Drivers d ON o.driver_id = d.driver_id JOIN OrderStatus os ON os.order_id = o.order_id GROUP BY month, d.driver_id, d.name ORDER BY month, total_orders DESC"
func (r *CustomerRepositoryDB) GetMonthlyTopDrivers() (interface{}, error) {
	query := "SELECT DATE_TRUNC('month', o.order_date) AS month, d.driver_id, d.name, COUNT(CASE WHEN os.status = 'completed' THEN o.order_id END) AS total_success_orders, COUNT(CASE WHEN os.status = 'cancelled' THEN o.order_id END) AS total_cancelled_orders,COUNT(o.order_id) AS total_orders FROM Orders o JOIN Drivers d ON o.driver_id = d.driver_id JOIN OrderStatus os ON os.order_id = o.order_id GROUP BY month, d.driver_id, d.name ORDER BY month, total_orders DESC"

	rows, err := r.DB.Query(query)
	if err != nil {
		return "", err
	}

	var GetMonthlyTopDrivers []struct {
		Month                  string
		Driver_id              int
		Name                   string
		Total_success_orders   int
		Total_cancelled_orders int
		Total_orders           int
	}

	for rows.Next() {
		MonthlyOrder := struct {
			Month                  string
			Driver_id              int
			Name                   string
			Total_success_orders   int
			Total_cancelled_orders int
			Total_orders           int
		}{}

		err := rows.Scan(&MonthlyOrder.Month, &MonthlyOrder.Driver_id, &MonthlyOrder.Name, &MonthlyOrder.Total_success_orders, &MonthlyOrder.Total_cancelled_orders, &MonthlyOrder.Total_orders)
		if err != nil {
			log.Fatal(err)
		}
		GetMonthlyTopDrivers = append(GetMonthlyTopDrivers, MonthlyOrder)
	}
	data := utils.ConvertSliceToMap(GetMonthlyTopDrivers)
	keys := utils.GetStructKeys(GetMonthlyTopDrivers[0], []string{})
	utils.DisplayData("Data Top Driver Tiap Bulan", data, keys)
	return GetMonthlyTopDrivers, nil
}
