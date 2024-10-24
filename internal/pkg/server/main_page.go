package server

import (
	"html/template"
	"log"
	"net/http"

	"zlatoivan_ru/internal/pkg/color"
)

type ViewData struct {
	Title   string
	Message string
}

func MainPage(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("static/template/index.html")
	err := t.Execute(w, "")
	if err != nil {
		log.Printf("t.Execute: %v", err)
	}
	log.Printf(color.Green("main page"))
}
