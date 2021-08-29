package router

import (
	"github.com/Budi721/homework_sql/handler"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func NewRouter(movieHandler handler.MovieHandler) chi.Router {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Post("/movies", movieHandler.Create)
	r.Get("/movies/{slug}", movieHandler.FindBySlug)
	return r
}