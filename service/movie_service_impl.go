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

func (m MovieServiceImpl) Update() {
	panic("implement me")
}

func (m MovieServiceImpl) Delete() {
	panic("implement me")
}

func (m MovieServiceImpl) FindBySlug(movieSlug string) web.MovieResponse {
	movie := m.MovieRepository.FindBySlug(m.DB, movieSlug)
	return helper.ToMovieResponse(movie)
}
