package orders

import (
	pagehandler "latihan/controller/pageHandler"
	"net/http"
)

func (o *Orders) AddOrderHandler(w http.ResponseWriter, r *http.Request) {
	pagehandler.RenderTemplate(w, "place-order.html", nil)
}
