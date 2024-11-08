package bookstore

import (
	"latihan/controller"
	"latihan/model/books"
	"latihan/service"
	"net/http"
	"strconv"
)

func AddBookHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		controller.RenderTemplate(w, "add-book.html", nil)
		return
	}

	if r.Method == "POST" {
		price, err := strconv.Atoi(r.FormValue("price"))
		if err != nil {
			controller.ErrorPage(w, err.Error())
			return
		}

		diskon := r.FormValue("discount")
		var discount float64
		if diskon != "" {
			diskon, err := strconv.ParseFloat(r.FormValue("discount"), 64)
			if err != nil {
				controller.ErrorPage(w, err.Error())
				return
			}

			discount = diskon
		} else {
			discount = 0
		}

		var data = books.Book{
			ID:       r.FormValue("bookID"),
			Name:     r.FormValue("bookName"),
			Type:     r.FormValue("bookType"),
			Author:   r.FormValue("author"),
			Price:    price,
			Discount: discount,
		}

		err = service.ServiceF.AddBookDataService(data)
		if err != nil {
			controller.ErrorPage(w, err.Error())
			return
		}

		http.Redirect(w, r, "/api/book-list", http.StatusSeeOther)
	}
}
