package service

import (
	"github.com/Budi721/homework_sql/helper"
	"github.com/Budi721/homework_sql/model/domain"
	"github.com/Budi721/homework_sql/model/web"
	"github.com/Budi721/homework_sql/repository"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
	"log"
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

func (m MovieServiceImpl) Create(request web.MovieRequest) (web.MovieResponse, error){
	err := m.Validate.Struct(request)
	if err != nil {
		return web.MovieResponse{}, err
	}

	var movie domain.Movie

	movie = domain.Movie{
		Title:       request.Title,
		Slug:        request.Slug,
		Description: request.Description,
		Duration:    request.Duration,
		Image:       request.Image,
	}

	movieInserted, err := m.MovieRepository.Save(m.DB, movie)
	if err != nil {
		return web.MovieResponse{}, err
	}

	return helper.ToMovieResponse(movieInserted), nil
}

func (m MovieServiceImpl) Update(request web.MovieRequest, movieSlug string) (web.MovieResponse, error) {
	err := m.Validate.Struct(request)
	if err != nil {
		return web.MovieResponse{}, err
	}

	movie := domain.Movie{
		Title:       request.Title,
		Slug:        request.Slug,
		Description: request.Description,
		Duration:    request.Duration,
		Image:       request.Image,
	}

	movieUpdated, err := m.MovieRepository.Update(m.DB, movie, movieSlug)
	if err != nil {
		return web.MovieResponse{}, err
	}

	return helper.ToMovieResponse(movieUpdated), nil
}

func (m MovieServiceImpl) Delete(movieSlug string) error {
	err := m.MovieRepository.Delete(m.DB, movieSlug)
	log.Println(err)
	if err != nil {
		return err
	}
	return nil
}

func (m MovieServiceImpl) FindBySlug(movieSlug string) (web.MovieResponse, error) {
	movie, err := m.MovieRepository.FindBySlug(m.DB, movieSlug)
	if err != nil {
		return web.MovieResponse{}, err
	}
	return helper.ToMovieResponse(movie), nil
}
