package middleware

import (
	"net/http"

	"github.com/zlatoivan/zlatoivan_ru/internal/pkg/metric"
)

// Metric - считает метрики
func Metric(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		if req.Method == http.MethodGet {
			metric.GetRequestsCounter.Inc()
		}
		next.ServeHTTP(w, req)
	})
}
