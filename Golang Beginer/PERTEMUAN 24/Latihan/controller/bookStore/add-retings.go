package bookstore

import (
	pagehandler "latihan/controller/pageHandler"
	"net/http"
)

func (b *Books) SubmitReviewHandler(w http.ResponseWriter, r *http.Request) {
	pagehandler.RenderTemplate(w, "submit-review.html", nil)
}
