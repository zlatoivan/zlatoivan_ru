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
	statusCode int
}

// WriteHeader overrides WriteHeader to capture status code
func (crw *customResponseWriter) WriteHeader(code int) {
	crw.statusCode = code
	crw.ResponseWriter.WriteHeader(code)
}

// Status returns the captured statusCode code
func (crw *customResponseWriter) Status() int {
	return crw.statusCode
}

func getReqTime(w http.ResponseWriter, req *http.Request, next http.Handler) (string, string) {
	//body, err := io.ReadAll(req.Body)
	//if err != nil {
	//	log.Println(err)
	//}
	//defer req.Body.Close()
	//length := int64(len(body))
	//fmt.Printf("%v %d\n", body, length)

	customRespWriter := customResponseWriter{w, http.StatusOK}
	t1 := time.Now()
	next.ServeHTTP(&customRespWriter, req)
	reqTime := color.NGreen(time.Since(t1).String())
	status := color.GetColoredStatus(customRespWriter.Status())
	return reqTime, status
}

// RequestLogger prints information about the request
func RequestLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		quote := color.NCyan("\"")
		method := color.Magenta(req.Method)
		proto := req.Header.Get("X-Forwarded-Proto")
		url := color.NCyan(fmt.Sprintf("%s://%s%s %s", proto, req.Host, req.URL, req.Proto))
		ip := req.Header.Get("X-Forwarded-For")
		reqTime, status := getReqTime(w, req, next)

		logs := fmt.Sprintf("%s%s %s%s from %s - %s in %s", quote, method, url, quote, ip, status, reqTime)
		log.Print(logs)
	})
}
