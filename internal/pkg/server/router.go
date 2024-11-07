package server

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/zlatoivan/zlatoivan_ru/internal/pkg/server/handlers"
	mw "github.com/zlatoivan/zlatoivan_ru/internal/pkg/server/middleware"
)

func createRouter() *chi.Mux {
	r := chi.NewRouter()

	r.Use(
		mw.RequestLogger,
		mw.StaticFileLoader,
		mw.Metric,
		middleware.Recoverer,
	)

	r.Handle("/metrics", promhttp.Handler())

	r.NotFound(handlers.PageNotFound)

	r.Get("/favicon.ico", handlers.FaviconIco)
	r.Get("/", handlers.MainPage)
	r.Get("/donut", handlers.Donut)

	return r
}
