package bookstore

import (
	"latihan/controller"
	"latihan/model/books"
	"latihan/service"
	"net/http"
)

func BookListHandler(w http.ResponseWriter, r *http.Request) {
	var data []books.Book
	err := service.ServiceF.GetBookDataService(&data)
	if err != nil {
		controller.ErrorPage(w, err.Error())
		return
	}

	result := map[string]interface{}{
		"Books": data,
	}
	controller.RenderTemplate(w, "book-list.html", result)
}
