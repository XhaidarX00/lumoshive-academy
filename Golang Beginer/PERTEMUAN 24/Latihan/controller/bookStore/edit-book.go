package bookstore

import (
	pagehandler "latihan/controller/pageHandler"
	"latihan/model/books"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
)

func (b *Books) EditBookHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		idString := chi.URLParam(r, "bookID")
		data := map[string]string{
			"ID": idString,
		}
		pagehandler.RenderTemplate(w, "edit-book.html", data)
		return
	}

	if r.Method == "POST" {
		idString := chi.URLParam(r, "bookID")
		price, err := strconv.Atoi(r.FormValue("price"))
		if err != nil {
			b.logger.Error("Error editbookshandler", zap.Error(err))
			pagehandler.ErrorPage(w, err.Error())
			return
		}

		diskon, err := strconv.ParseFloat(r.FormValue("discount"), 64)
		if err != nil {
			b.logger.Error("Error editbookshandler", zap.Error(err))
			pagehandler.ErrorPage(w, err.Error())
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

		err = b.Service.EditBookDataService(data)
		if err != nil {
			b.logger.Error("Error editbookshandler", zap.Error(err))
			pagehandler.ErrorPage(w, err.Error())
			return
		}

		http.Redirect(w, r, "/api/book-list", http.StatusSeeOther)
	}
}
