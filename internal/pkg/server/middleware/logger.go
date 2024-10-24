package middleware

import (
	"fmt"
	"log"
	"net/http"
	"time"
	"zlatoivan_ru/internal/pkg/color"
)

type statusRecorder struct {
	http.ResponseWriter
	status int
}

func (rec *statusRecorder) WriteHeader(code int) {
	rec.status = code
	rec.ResponseWriter.WriteHeader(code)
}

func ReqLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		quoteCyan := color.NCyan("\"")
		method := color.Magenta(fmt.Sprintf("%s", req.Method))
		scheme := "http"
		if req.TLS != nil {
			scheme = "https"
		}
		url := color.NCyan(fmt.Sprintf("%s://%s%s %s", scheme, req.Host, req.URL, req.Proto))

		rec := statusRecorder{w, 200}
		t1 := time.Now()
		next.ServeHTTP(&rec, req)
		reqTime := color.NGreen(time.Since(t1).String())

		coloredStatus := color.GetColoredStatus(rec.status)

		logs := fmt.Sprintf("%s%s %s%s - %s in %s", quoteCyan, method, url, quoteCyan, coloredStatus, reqTime)
		log.Print(logs)
	})
}
