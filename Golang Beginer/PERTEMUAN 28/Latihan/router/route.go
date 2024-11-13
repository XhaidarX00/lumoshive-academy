package router

import (
	"latihan/controller"
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
)

func InitRoute(ch *controller.Travel, logger *zap.Logger) {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/api/", ch.TravelController)

	r.MethodNotAllowed(MethodNotAllowedHandler)

	logger.Info("server started on http://localhost:8080")
	http.ListenAndServe(":8080", r)
}

func MethodNotAllowedHandler(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
}
