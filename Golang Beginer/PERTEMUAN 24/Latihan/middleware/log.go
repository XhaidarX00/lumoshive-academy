package middlewaree

import (
	"latihan/service"
	"net/http"
	"time"

	"go.uber.org/zap"
)

type Middlewaree struct {
	svc *service.Service
	log *zap.Logger
}

func NewMiddleware(log *zap.Logger, svc *service.Service) *Middlewaree {
	return &Middlewaree{
		svc: svc,
		log: log,
	}
}

func (middlewaree *Middlewaree) MiddlewareeLogger(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		handler.ServeHTTP(w, r)
		duration := time.Since(start)
		middlewaree.log.Info("Http Request",
			zap.String("url", r.URL.String()),
			zap.String("method", r.Method),
			zap.Duration("duration", duration),
		)
	})
}
