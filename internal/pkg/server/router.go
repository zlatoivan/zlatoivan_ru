package server

import (
	"zlatoivan_ru/internal/pkg/server/handlers"
	mw "zlatoivan_ru/internal/pkg/server/middleware"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (s Server) createRouter() *chi.Mux {
	r := chi.NewRouter()

	r.Use(
		mw.RequestLoggerMW,
		mw.StaticFileMW,
		middleware.Recoverer,
		middleware.Logger,
	)

	r.Get("/favicon.ico", handlers.FaviconIco)

	r.Get("/", handlers.MainPage)
	r.Get("/donut", handlers.Donut)

	return r
}
