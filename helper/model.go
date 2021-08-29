package helper

import (
	"github.com/Budi721/homework_sql/model/domain"
	"github.com/Budi721/homework_sql/model/web"
)

func ToMovieResponse(movie domain.Movie) web.MovieResponse {
	return web.MovieResponse{
		Id: movie.Id,
		MovieRequest: web.MovieRequest{
			Title:       movie.Title,
			Slug:        movie.Slug,
			Description: movie.Description,
			Duration:    movie.Duration,
			Image:       movie.Image,
		},
	}
}
