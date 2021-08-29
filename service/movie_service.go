package service

import "github.com/Budi721/homework_sql/model/web"

type MovieService interface {
	Create(request web.MovieRequest) web.MovieResponse
	Update()
	Delete()
	FindBySlug(movieSlug string) web.MovieResponse
}
