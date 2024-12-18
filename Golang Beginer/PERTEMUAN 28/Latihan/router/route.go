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
	r.Get("/api/detail", ch.PlaceDetailController)
	r.Post("/api/add-transaction", ch.AddTransactionController)
	r.Get("/api/location", ch.GetLocationByIdController)
	r.Get("/api/tour-plan", ch.GetTourPLanByIdController)

	r.MethodNotAllowed(MethodNotAllowedHandler)

	logger.Info("server started on http://localhost:8080")
	http.ListenAndServe(":8080", r)
}

func MethodNotAllowedHandler(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
}
