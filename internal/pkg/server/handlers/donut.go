package handlers

import (
	"html/template"
	"log"
	"net/http"

	"zlatoivan_ru/internal/pkg/color"
)

func Donut(w http.ResponseWriter, _ *http.Request) {
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
