package handlers

import "net/http"

// FaviconIco - подгружает логотип для вкладки
func FaviconIco(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "static/images/favicon.ico")
}
