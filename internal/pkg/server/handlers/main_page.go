package handlers

import (
	"html/template"
	"log"
	"net/http"

	"github.com/zlatoivan/zlatoivan_ru/internal/pkg/color"
)

// MainPage - главная страница
func MainPage(w http.ResponseWriter, _ *http.Request) {
	t, err := template.ParseFiles("static/template/index.html")
	if err != nil {
		log.Printf("template.ParseFiles: %v", err)
		http.Error(w, "Internal server Error", http.StatusInternalServerError)
		return
	}

	err = t.Execute(w, "")
	if err != nil {
		log.Printf("t.Execute: %v", err)
		http.Error(w, "Internal server Error", http.StatusInternalServerError)
		return
	}

	log.Print(color.Green("main page"))
}
