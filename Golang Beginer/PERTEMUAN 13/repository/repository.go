package repository

import "main/model"

type CustomerRepository interface {
	CreateCustomers(customer *model.Customers) (uint16, error)
	CreateDrivers(customer *model.Drivers) (uint16, error)
	CreateOrders(customer *model.Orders) (uint16, error)
	ViewMonthlyOrders() (interface{}, error)
	ViewMonthlyCustomerOrders() (interface{}, error)
	GetMonthlyCustomerOrders() (interface{}, error)
	ViewOrderPeakTimes() (interface{}, error)
	ViewCustomersSession() (interface{}, error)
	GetMonthlyTopDrivers() (interface{}, error)
}
