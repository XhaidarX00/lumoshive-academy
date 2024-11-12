package bookstore

import (
	pagehandler "latihan/controller/pageHandler"
	"latihan/model/books"
	"latihan/service"
	"net/http"
	"strconv"

	"go.uber.org/zap"
)

type Books struct {
	Service *service.Service
	logger  *zap.Logger
}

func NewBooksHandelr(serv *service.Service, log *zap.Logger) *Books {
	return &Books{
		Service: serv,
		logger:  log,
	}
}

func (b *Books) AddBookHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		pagehandler.RenderTemplate(w, "add-book.html", nil)
		return
	}

	if r.Method == "POST" {
		price, err := strconv.Atoi(r.FormValue("price"))
		if err != nil {
			b.logger.Error("Error addbookshandler", zap.Error(err))
			pagehandler.ErrorPage(w, err.Error())
			return
		}

		diskon := r.FormValue("discount")
		var discount float64
		if diskon != "" {
			diskon, err := strconv.ParseFloat(r.FormValue("discount"), 64)
			if err != nil {
				b.logger.Error("Error addbookshandler", zap.Error(err))
				pagehandler.ErrorPage(w, err.Error())
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

		err = b.Service.AddBookDataService(data)
		if err != nil {
			b.logger.Error("Error addbookshandler", zap.Error(err))
			pagehandler.ErrorPage(w, err.Error())
			return
		}

		http.Redirect(w, r, "/api/book-list", http.StatusSeeOther)
	}
}
