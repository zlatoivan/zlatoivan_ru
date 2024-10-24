package server

import (
	"html/template"
	"log"
	"net/http"

	"zlatoivan_ru/internal/pkg/color"
)

func Donut(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("static/template/donut.html")
	err := t.Execute(w, "")
	if err != nil {
		log.Printf("t.Execute: %v", err)
	}
	log.Printf(color.Green("donut"))
}
