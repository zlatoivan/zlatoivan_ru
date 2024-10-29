package handlers

import "net/http"

func FaviconIco(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "static/images/favicon.ico")
	w.WriteHeader(http.StatusOK)
}
