package middleware

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/zlatoivan/zlatoivan_ru/internal/pkg/color"
)

func getProto(req *http.Request) string {
	scheme := "http"
	if req.TLS != nil || req.Header.Get("X-Forwarded-Proto") == "https" {
		scheme = "https"
	}
	return scheme
}

func getIP(req *http.Request) string {
	ip := req.Header.Get("X-Forwarded-For")
	if ip == "" {
		ip = req.RemoteAddr
	}
	return ip
}

// RequestLogger - логгирует информацию о запросе
func RequestLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		quote := color.NCyan("\"")
		method := color.BNMagenta(req.Method)
		proto := getProto(req)
		url := fmt.Sprintf("%s://%s%s", proto, req.Host, req.URL)
		protoFull := color.NCyan(req.Proto)

		ip := getIP(req)

		customRespWriter := NewWrapResponseWriter(w)
		timeStart := time.Now()
		next.ServeHTTP(customRespWriter, req)
		reqTime := color.NGreen(time.Since(timeStart).String())
		status := color.Status(customRespWriter.Status())
		bytes := color.BNBlue(strconv.Itoa(customRespWriter.BytesWritten()) + "B")

		// "GET https://zlatoivan.ru/ HTTP/1.1" from 46.138.82.38 - 200 9078B in 1.342625ms
		logs := fmt.Sprintf("%s%s %s %s%s from %s - %s %s in %s", quote, method, url, protoFull, quote, ip, status, bytes, reqTime)
		log.Print(logs)
	})
}
