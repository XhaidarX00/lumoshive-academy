package bookstore

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
)

func (b *Books) DeleteBookHandler(w http.ResponseWriter, r *http.Request) {
	idString := chi.URLParam(r, "bookID")
	if err := b.Service.DeleteBookService(idString); err != nil {
		b.logger.Error("Error deletebookhandler", zap.Error(err))
		return
	}

	http.Redirect(w, r, "/api/book-list", http.StatusSeeOther)
}
