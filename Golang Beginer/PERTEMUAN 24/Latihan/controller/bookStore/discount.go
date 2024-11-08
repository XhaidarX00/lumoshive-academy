package bookstore

import (
	"latihan/controller"
	"net/http"
)

func DiscountBookHandler(w http.ResponseWriter, r *http.Request) {
	controller.RenderTemplate(w, "discount-book.html", nil)
}
