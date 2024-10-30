package main

import (
	"log"
	"main/database"
	"main/service"
)

func main() {

	db, err := database.InitDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// newCustomer := model.Customers{
	// 	Name:         "aldydiana",
	// 	Email:        "aldydiana@example.com",
	// 	Phone_number: "081234567815",
	// 	Created_at:   time.Now(),
	// }
	// service.InputDataCustomer(db, newCustomer)

	// newDriver := model.Drivers{
	// 	Name:         "Mimin",
	// 	Phone_number: "0897898774101",
	// 	Vehicle_type: "Car",
	// 	Created_at:   time.Now(),
	// }

	// service.InputDataDriver(db, newDriver)

	// newOrder := model.Orders{
	// 	Customer_id:      2,
	// 	Driver_id:        2,
	// 	Pickup_location:  2,
	// 	Dropoff_location: 2,
	// 	Total_fare:       75000,
	// 	Order_date:       time.Now(),
	// }

	// service.InputDataOrders(db, newOrder)

	// service.ViewMonthlyOrders(db)
	// service.ViewMonthlyCustomerOrders(db)
	// service.GetMonthlyCustomerOrders(db)
	// service.ViewOrderPeakTimes(db)
	// service.ViewCustomersSession(db)
	service.GetMonthlyTopDrivers(db)
}
