package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		// Устанавливаем заголовок ответа
		w.Header().Set("Content-Type", "text/plain")

		// Выводим все заголовки запроса в консоль
		fmt.Println("Request Headers:")
		for name, values := range r.Header {
			for _, value := range values {
				fmt.Printf("%s: %s\n", name, value)
			}
		}

		// Ответ клиенту
		fmt.Fprintln(w, "Headers received. Check the console for details.")

		fmt.Println()
	})

	fmt.Println("Сервер запущен на http://localhost:7070")
	http.ListenAndServe(":7070", r)
}
