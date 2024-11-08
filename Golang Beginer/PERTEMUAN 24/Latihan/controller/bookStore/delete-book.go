package bookstore

import (
	"latihan/service"
	"net/http"

	"github.com/go-chi/chi"
)

func DeleteBookHandler(w http.ResponseWriter, r *http.Request) {
	idString := chi.URLParam(r, "bookID")
	if err := service.ServiceF.DeleteBookService(idString); err != nil {
		panic(err)
	}

	http.Redirect(w, r, "/api/book-list", http.StatusSeeOther)

}
