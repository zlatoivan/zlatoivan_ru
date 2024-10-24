package server

import (
	mw "zlatoivan_ru/internal/pkg/server/middleware"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (s Server) createRouter() *chi.Mux {
	r := chi.NewRouter()
	r.Use(
		//middleware.RequestID,
		//middleware.Logger,
		mw.ReqLogger,
		middleware.Recoverer,
	)

	r.Get("/*", LoadStatic)
	r.Get("/favicon.ico", FaviconIco)

	r.Get("/", MainPage)

	return r
}
