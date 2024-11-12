package orders

import (
	pagehandler "latihan/controller/pageHandler"
	"latihan/model/orders"
	"net/http"

	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
)

func (o *Orders) OrderDetailHandler(w http.ResponseWriter, r *http.Request) {
	var data orders.Order
	idString := chi.URLParam(r, "orderID")
	if idString == "" {
		o.logger.Error("Id Order Tidak ditemukan")
		pagehandler.ErrorPage(w, "Id Order Tidak ditemukan")
		return
	}
	data.ID = idString
	err := o.Service.GetOrderDetailService(&data)
	if err != nil {
		o.logger.Error("Error Service :", zap.Error(err))
		pagehandler.ErrorPage(w, err.Error())
		return
	}

	result := map[string]interface{}{
		"Orders": data,
	}

	pagehandler.RenderTemplate(w, "order-detail.html", result)
}
