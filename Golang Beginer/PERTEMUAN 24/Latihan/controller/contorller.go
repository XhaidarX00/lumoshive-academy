package controller

import (
	bookstore "latihan/controller/bookStore"
	"latihan/controller/orders"
	payments "latihan/controller/payment"
	usershandler "latihan/controller/users_handler"
	"latihan/service"

	"go.uber.org/zap"
)

type Controller struct {
	Payments     *payments.Payment
	Books        *bookstore.Books
	Orders       *orders.Orders
	Usershandler *usershandler.Auth
}

func NewController(svc *service.Service, logger *zap.Logger) *Controller {
	return &Controller{
		Payments:     payments.NewPaymenHandelr(svc, logger),
		Books:        bookstore.NewBooksHandelr(svc, logger),
		Orders:       orders.NewOrdersHandelr(svc, logger),
		Usershandler: usershandler.NewUserHandelr(svc, logger),
	}
}
