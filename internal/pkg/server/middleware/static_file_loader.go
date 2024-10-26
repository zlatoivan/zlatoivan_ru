package middleware

import (
	"net/http"
	"strings"
)

// StaticFileLoader serves static files from the specified directory.
func StaticFileLoader(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.Contains(r.URL.Path, "/static/") {
			fs := http.FileServer(http.Dir("static"))
			http.StripPrefix("/static/", fs).ServeHTTP(w, r)
			return
		}
		next.ServeHTTP(w, r)
	})
}
