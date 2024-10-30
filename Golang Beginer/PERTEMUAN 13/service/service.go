package service

import (
	"database/sql"
	"fmt"
	utils "main/Utils"
	"main/model"
	"main/repository"
)

func InputDataCustomer(db *sql.DB, customer model.Customers) {
	if customer.Name == "" {
		utils.ErrorMessage("name tidak boleh kosong")
		return
	}
	if customer.Email == "" {
		utils.ErrorMessage("email tidak boleh kosong")
		return
	}
	if customer.Phone_number == "" {
		utils.ErrorMessage("phone number tidak boleh kosong")
		return
	}

	customerRepo := repository.NewCustomRepository(db)

	newCustomer := model.Customers{
		Name:         customer.Name,
		Email:        customer.Email,
		Phone_number: customer.Phone_number,
		Created_at:   customer.Created_at,
	}

	result, err := customerRepo.CreateCustomers(&newCustomer)
	if err != nil {
		utils.ErrorMessage(err.Error())
		return
	}

	msg := fmt.Sprintf("Berhasil input data customer dengan id %v\n", result)
	utils.SuccesMessage(msg)
	return
}

func InputDataDriver(db *sql.DB, driver model.Drivers) {
	if driver.Name == "" {
		utils.ErrorMessage("name tidak boleh kosong")
		return
	}
	if driver.Phone_number == "" {
		utils.ErrorMessage("password tidak boleh kosong")
		return
	}

	customerRepo := repository.NewCustomRepository(db)

	newDriver := model.Drivers{
		Name:         driver.Name,
		Phone_number: driver.Phone_number,
		Vehicle_type: driver.Vehicle_type,
		Created_at:   driver.Created_at,
	}

	result, err := customerRepo.CreateDrivers(&newDriver)
	if err != nil {
		utils.ErrorMessage(err.Error())
		return
	}

	msg := fmt.Sprintf("Berhasil input data customer dengan id %v\n", result)
	utils.SuccesMessage(msg)
	return
}

func InputDataOrders(db *sql.DB, order model.Orders) {
	if order.Customer_id < 0 {
		utils.ErrorMessage("customer id invalid")
		return
	}
	if order.Driver_id < 0 {
		utils.ErrorMessage("driver id invalid")
		return
	}
	if order.Pickup_location < 0 {
		utils.ErrorMessage("pickup location invalid")
		return
	}

	if order.Dropoff_location < 0 {
		utils.ErrorMessage("dropoff location invalid")
		return
	}

	if order.Total_fare < 0 {
		utils.ErrorMessage("Total order tidak boleh dibawah 0")
		return
	}

	customerRepo := repository.NewCustomRepository(db)

	newOrder := model.Orders{
		Customer_id:      order.Customer_id,
		Driver_id:        order.Driver_id,
		Pickup_location:  order.Pickup_location,
		Dropoff_location: order.Dropoff_location,
		Total_fare:       order.Total_fare,
		Order_date:       order.Order_date,
	}

	result, err := customerRepo.CreateOrders(&newOrder)
	if err != nil {
		utils.ErrorMessage(err.Error())
		return
	}

	msg := fmt.Sprintf("Berhasil input data customer dengan id %v\n", result)
	utils.SuccesMessage(msg)
}

// -- View: Total order setiap bulan
func ViewMonthlyOrders(db *sql.DB) {
	customerRepo := repository.NewCustomRepository(db)

	result, err := customerRepo.ViewMonthlyOrders()
	if err != nil {
		utils.ErrorMessage(err.Error())
		return
	}

	msg := fmt.Sprintf("Result: %v\n", result)
	utils.SuccesMessage(msg)
}

// -- View: Customer yang sering order tiap bulan
func ViewMonthlyCustomerOrders(db *sql.DB) {
	customerRepo := repository.NewCustomRepository(db)

	result, err := customerRepo.ViewMonthlyCustomerOrders()
	if err != nil {
		utils.ErrorMessage(err.Error())
		return
	}

	msg := fmt.Sprintf("Result: %v\n", result)
	utils.SuccesMessage(msg)
}

// -- View: Daerah dengan banyak order
func GetMonthlyCustomerOrders(db *sql.DB) {
	customerRepo := repository.NewCustomRepository(db)

	result, err := customerRepo.GetMonthlyCustomerOrders()
	if err != nil {
		utils.ErrorMessage(err.Error())
		return
	}

	msg := fmt.Sprintf("Result: %v\n", result)
	utils.SuccesMessage(msg)
}

// -- View: Pukul berapa saja order yang ramai dan sepi
func ViewOrderPeakTimes(db *sql.DB) {
	customerRepo := repository.NewCustomRepository(db)

	result, err := customerRepo.ViewOrderPeakTimes()
	if err != nil {
		utils.ErrorMessage(err.Error())
		return
	}

	msg := fmt.Sprintf("Result: %v\n", result)
	utils.SuccesMessage(msg)
}

// -- View: Jumlah customer yang masih login dan belum logout
func ViewCustomersSession(db *sql.DB) {
	customerRepo := repository.NewCustomRepository(db)

	result, err := customerRepo.ViewCustomersSession()
	if err != nil {
		utils.ErrorMessage(err.Error())
		return
	}

	msg := fmt.Sprintf("Result: %v\n", result)
	utils.SuccesMessage(msg)
}

// -- View: Driver yang rajin mengambil order setiap bulan
func GetMonthlyTopDrivers(db *sql.DB) {
	customerRepo := repository.NewCustomRepository(db)

	result, err := customerRepo.GetMonthlyTopDrivers()
	if err != nil {
		utils.ErrorMessage(err.Error())
		return
	}

	msg := fmt.Sprintf("Result: %v\n", result)
	utils.SuccesMessage(msg)
}
