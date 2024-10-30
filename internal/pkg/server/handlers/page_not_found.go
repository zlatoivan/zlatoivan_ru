package handlers

import (
	"log"
	"net/http"

	"github.com/zlatoivan/zlatoivan_ru/internal/pkg/color"
)

// PageNotFound informs that the page is not found
func PageNotFound(w http.ResponseWriter, _ *http.Request) {
	log.Println(color.Green("page not found"))
	http.Error(w, "404 page not found", http.StatusNotFound)
}
