package orders

import (
	"latihan/controller"
	"net/http"
)

func OrderDetailHandler(w http.ResponseWriter, r *http.Request) {
	// var data []orders.Order
	// err := service.ServiceF.GetOrderDataService(&data)
	// if err != nil {
	// 	controller.ErrorPage(w, err.Error())
	// 	return
	// }

	// result := map[string]interface{}{
	// 	"Orders": data,
	// }
	controller.RenderTemplate(w, "order-detail.html", nil)
}
