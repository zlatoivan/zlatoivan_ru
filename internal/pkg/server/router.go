package server

import (
	"zlatoivan_ru/internal/pkg/server/handlers"
	mw "zlatoivan_ru/internal/pkg/server/middleware"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func (s Server) createRouter() *chi.Mux {
	r := chi.NewRouter()

	r.Use(
		middleware.Recoverer,
		mw.RequestLogger,
		mw.StaticFileLoader,
		mw.Metric,
	)

	r.Handle("/metrics", promhttp.Handler())

	r.NotFound(handlers.PageNotFound)

	r.Get("/favicon.ico", handlers.FaviconIco)
	r.Get("/", handlers.MainPage)
	r.Get("/donut", handlers.Donut)

	return r
}
