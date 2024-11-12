package bookstore

import (
	pagehandler "latihan/controller/pageHandler"
	"latihan/model/books"
	"net/http"

	"go.uber.org/zap"
)

func (b *Books) BookListHandler(w http.ResponseWriter, r *http.Request) {
	var data []books.Book
	err := b.Service.GetBookDataService(&data)
	if err != nil {
		b.logger.Error("Error bookslisthandler", zap.Error(err))
		pagehandler.ErrorPage(w, err.Error())
		return
	}

	result := map[string]interface{}{
		"Books": data,
	}
	pagehandler.RenderTemplate(w, "book-list.html", result)
}
