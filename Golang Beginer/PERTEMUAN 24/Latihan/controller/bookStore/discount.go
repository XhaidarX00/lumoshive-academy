package bookstore

import (
	pagehandler "latihan/controller/pageHandler"
	"net/http"
)

func (b *Books) DiscountBookHandler(w http.ResponseWriter, r *http.Request) {
	pagehandler.RenderTemplate(w, "discount-book.html", nil)
}
