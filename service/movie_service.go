package service

import "github.com/Budi721/homework_sql/model/web"

type MovieService interface {
	Create(request web.MovieRequest) web.MovieResponse
	Update(request web.MovieRequest, movieSlug string) web.MovieResponse
	Delete(movieSlug string)
	FindBySlug(movieSlug string) web.MovieResponse
}
