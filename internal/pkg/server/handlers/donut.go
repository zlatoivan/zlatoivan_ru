package handlers

import (
	"html/template"
	"log"
	"net/http"
	"strings"

	"zlatoivan_ru/internal/pkg/color"
	"zlatoivan_ru/internal/pkg/donut"
)

func Donut(w http.ResponseWriter, r *http.Request) {
	userAgent := r.Header.Get("User-Agent")
	if strings.Contains(userAgent, "curl") {
		err := donut.SendDonutToConsole(w)
		if err != nil {
			log.Printf("sendDonutToConsole: %w", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
		return
	}

	t, err := template.ParseFiles("static/template/donut.html")
	if err != nil {
		log.Printf("template.ParseFiles: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = t.Execute(w, "")
	if err != nil {
		log.Printf("t.Execute: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	log.Print(color.Green("donut"))
}
