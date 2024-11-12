package orders

import (
	pagehandler "latihan/controller/pageHandler"
	"latihan/model/orders"
	"latihan/service"
	"net/http"

	"go.uber.org/zap"
)

type Orders struct {
	Service *service.Service
	logger  *zap.Logger
}

func NewOrdersHandelr(serv *service.Service, log *zap.Logger) *Orders {
	return &Orders{
		Service: serv,
		logger:  log,
	}
}

func (o *Orders) OrderListHandler(w http.ResponseWriter, r *http.Request) {
	var data []orders.Order
	err := o.Service.GetOrderDataService(&data)
	if err != nil {
		o.logger.Error("Error ", zap.Error(err))
		pagehandler.ErrorPage(w, err.Error())
		return
	}

	result := map[string]interface{}{
		"Orders": data,
	}

	pagehandler.RenderTemplate(w, "order-list.html", result)
}
