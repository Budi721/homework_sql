package service

import "github.com/Budi721/homework_sql/model/web"

type MovieService interface {
	Create(request web.MovieRequest) (web.MovieResponse, error)
	Update(request web.MovieRequest, movieSlug string) (web.MovieResponse, error)
	Delete(movieSlug string) error
	FindBySlug(movieSlug string) (web.MovieResponse, error)
}
