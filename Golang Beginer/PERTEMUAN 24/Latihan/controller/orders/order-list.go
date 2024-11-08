package orders

import (
	"fmt"
	"latihan/controller"
	"latihan/model/orders"
	"latihan/service"
	"net/http"
)

func OrderListHandler(w http.ResponseWriter, r *http.Request) {
	var data []orders.Order
	err := service.ServiceF.GetOrderDataService(&data)
	if err != nil {
		controller.ErrorPage(w, err.Error())
		return
	}

	result := map[string]interface{}{
		"Orders": data,
	}

	fmt.Println("%v\n", data)
	controller.RenderTemplate(w, "order-list.html", result)
}
