package bookstore

import (
	"fmt"
	"latihan/controller"
	"latihan/model/books"
	"latihan/service"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

func EditBookHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		idString := chi.URLParam(r, "bookID")
		fmt.Println(idString)
		data := map[string]string{
			"ID": idString,
		}
		controller.RenderTemplate(w, "edit-book.html", data)
		return
	}

	if r.Method == "POST" {
		idString := chi.URLParam(r, "bookID")
		fmt.Println(idString)
		price, err := strconv.Atoi(r.FormValue("price"))
		if err != nil {
			controller.ErrorPage(w, err.Error())
			return
		}

		diskon, err := strconv.ParseFloat(r.FormValue("discount"), 64)
		if err != nil {
			controller.ErrorPage(w, err.Error())
			return
		}

		var data = books.Book{
			ID:       idString,
			Name:     r.FormValue("bookName"),
			Type:     r.FormValue("bookType"),
			Author:   r.FormValue("author"),
			Price:    price,
			Discount: diskon,
		}

		err = service.ServiceF.EditBookDataService(data)
		if err != nil {
			controller.ErrorPage(w, err.Error())
			return
		}

		http.Redirect(w, r, "/api/book-list", http.StatusSeeOther)
	}
}
