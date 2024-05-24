package server

import (
	"html/template"
	"log"
	"net/http"
)

type ViewData struct {
	Title   string
	Message string
}

func MainPage(w http.ResponseWriter, r *http.Request) {
	data := ViewData{
		Title:   "World Cup",
		Message: "FIFA will never regret it",
	}

	t, _ := template.ParseFiles("templates/index.html")
	err := t.Execute(w, data)
	if err != nil {
		log.Printf("t.Execute: %v", err)
	}
}
