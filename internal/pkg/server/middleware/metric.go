package middleware

import (
	"net/http"
	"zlatoivan_ru/internal/pkg/metric"
)

func Metric(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		if req.Method == http.MethodConnect {
			metric.GetRequestsCounter.Inc()
		}
		next.ServeHTTP(w, req)
	})
}
