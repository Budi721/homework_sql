package service

import (
	"github.com/Budi721/homework_sql/helper"
	"github.com/Budi721/homework_sql/model/domain"
	"github.com/Budi721/homework_sql/model/web"
	"github.com/Budi721/homework_sql/repository"
	"github.com/go-playground/validator"
	"gorm.io/gorm"
)

type MovieServiceImpl struct {
	MovieRepository repository.MovieRepository
	DB              *gorm.DB
	Validate        *validator.Validate
}

func NewMovieService(movieRepository repository.MovieRepository, db *gorm.DB, validate *validator.Validate) MovieService {
	return &MovieServiceImpl{
		MovieRepository: movieRepository,
		DB:              db,
		Validate:        validate,
	}
}

func (m MovieServiceImpl) Create(request web.MovieRequest) web.MovieResponse {
	db := m.DB

	movie := domain.Movie{
		Title:       request.Title,
		Slug:        request.Slug,
		Description: request.Description,
		Duration:    request.Duration,
		Image:       request.Image,
	}

	movieInserted := m.MovieRepository.Save(db, movie)
	return helper.ToMovieResponse(movieInserted)
}

func (m MovieServiceImpl) Update(request web.MovieRequest, movieSlug string) web.MovieResponse {
	movie := domain.Movie{
		Title:       request.Title,
		Slug:        request.Slug,
		Description: request.Description,
		Duration:    request.Duration,
		Image:       request.Image,
	}

	movieUpdated := m.MovieRepository.Update(m.DB, movie, movieSlug)
	return helper.ToMovieResponse(movieUpdated)
}

func (m MovieServiceImpl) Delete(movieSlug string) {
	m.MovieRepository.Delete(m.DB, movieSlug)
}

func (m MovieServiceImpl) FindBySlug(movieSlug string) web.MovieResponse {
	movie := m.MovieRepository.FindBySlug(m.DB, movieSlug)
	return helper.ToMovieResponse(movie)
}
