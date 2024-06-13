//go:generate minimock -i Producer -o mock/data_logger_mock.go -p mock -g

package middleware

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"zlatoivan_ru/internal/utils"
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
		logs := utils.Color("\"", "nCyan")
		logs += utils.Color(fmt.Sprintf("%s ", req.Method), "magenta")
		scheme := "http"
		if req.TLS != nil {
			scheme = "https"
		}
		logs += utils.Color(fmt.Sprintf("%s://%s%s %s\" ", scheme, req.Host, req.URL, req.Proto), "nCyan")

		rec := statusRecorder{w, 200}

		t1 := time.Now()

		next.ServeHTTP(&rec, req)

		var color string
		switch {
		case rec.status < 200:
			color = "blue"
		case rec.status < 300:
			color = "green"
		case rec.status < 400:
			color = "cyan"
		case rec.status < 500:
			color = "yellow"
		default:
			color = "red"
		}
		logs += "- " + utils.Color(fmt.Sprintf("%d ", rec.status), color) + "in "
		logs += utils.Color(time.Since(t1).String(), "nGreen")
		logs += "\n"
		log.Print(logs)
	})
}
