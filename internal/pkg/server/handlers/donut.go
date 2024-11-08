package handlers

import (
	"html/template"
	"log"
	"net/http"
	"strings"

	"github.com/zlatoivan/zlatoivan_ru/internal/pkg/color"
	"github.com/zlatoivan/zlatoivan_ru/internal/pkg/donut"
)

// Donut - высчитывает пончик по формуле и выводит
func Donut(w http.ResponseWriter, req *http.Request) {
	userAgent := req.Header.Get("User-Agent")
	if strings.Contains(userAgent, "curl") {
		err := donut.SendDonutToConsole(w, req)
		if err != nil {
			log.Printf("sendDonutToConsole: %v", err)
			http.Error(w, "Internal HTTPServer Error", http.StatusInternalServerError)
		}
		log.Print(color.Green("curl donut"))
		return
	}

	t, err := template.ParseFiles("static/template/donut.html")
	if err != nil {
		log.Printf("template.ParseFiles: %v", err)
		http.Error(w, "Internal HTTPServer Error", http.StatusInternalServerError)
		return
	}

	err = t.Execute(w, "")
	if err != nil {
		log.Printf("t.Execute: %v", err)
		http.Error(w, "Internal HTTPServer Error", http.StatusInternalServerError)
		return
	}

	log.Print(color.Green("donut"))
}
