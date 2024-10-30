package middleware

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/zlatoivan/zlatoivan_ru/internal/pkg/color"
)

type customResponseWriter struct {
	http.ResponseWriter
	statusCode int
	bytes      int
}

// WriteHeader overrides WriteHeader to capture status code
func (crw *customResponseWriter) WriteHeader(status int) {
	crw.statusCode = status
	crw.ResponseWriter.WriteHeader(status)
}

// Write captures the number of bytes written to the response.
func (crw *customResponseWriter) Write(b []byte) (int, error) {
	n, err := crw.ResponseWriter.Write(b)
	crw.bytes += n
	return n, err
}

// Status returns the captured statusCode code
func (crw *customResponseWriter) Status() int {
	return crw.statusCode
}

func getReqResults(w http.ResponseWriter, req *http.Request, next http.Handler) (string, string, string) {
	customRespWriter := customResponseWriter{w, http.StatusOK, 0}
	t1 := time.Now()
	next.ServeHTTP(&customRespWriter, req)
	reqTime := color.NGreen(time.Since(t1).String())
	status := color.GetColoredStatus(customRespWriter.Status())
	bytes := color.BNBlue(strconv.Itoa(customRespWriter.bytes) + "B")
	return reqTime, status, bytes
}

func getProto(req *http.Request) string {
	scheme := "http"
	if req.TLS != nil {
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

// RequestLogger prints information about the request
func RequestLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		quote := color.NCyan("\"")
		method := color.BNMagenta(req.Method)
		proto := getProto(req)
		url := color.NCyan(fmt.Sprintf("%s://%s%s %s", proto, req.Host, req.URL, req.Proto))
		ip := getIP(req)
		reqTime, status, respBytes := getReqResults(w, req, next)

		logs := fmt.Sprintf("%s%s %s%s from %s - %s %s in %s", quote, method, url, quote, ip, status, respBytes, reqTime)
		log.Print(logs)
	})
}
