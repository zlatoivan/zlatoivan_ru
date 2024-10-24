package middleware

import (
	"fmt"
	"log"
	"net/http"
	"time"
	"zlatoivan_ru/internal/pkg/color"
)

type customResponseWriter struct {
	http.ResponseWriter
	status int
}

func (crw *customResponseWriter) WriteHeader(code int) {
	crw.status = code
	crw.ResponseWriter.WriteHeader(code)
}

// RequestLoggerMW prints information about the request
func RequestLoggerMW(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		quote := color.NCyan("\"")
		method := color.Magenta(req.Method)
		scheme := "http"
		if req.TLS != nil {
			scheme = "https"
		}
		url := color.NCyan(fmt.Sprintf("%s://%s%s %s", scheme, req.Host, req.URL, req.Proto))

		customRespWriter := customResponseWriter{w, http.StatusOK}
		t1 := time.Now()
		next.ServeHTTP(&customRespWriter, req)
		reqTime := color.NGreen(time.Since(t1).String())

		status := color.GetColoredStatus(customRespWriter.status)

		logs := fmt.Sprintf("%s%s %s%s - %s in %s", quote, method, url, quote, status, reqTime)
		log.Print(logs)
	})
}
