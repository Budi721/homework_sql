package handler

import (
	"github.com/Budi721/homework_sql/helper"
	"github.com/Budi721/homework_sql/model/web"
	"github.com/Budi721/homework_sql/service"
	"github.com/go-chi/chi"
	"net/http"
)

type MovieHandlerImpl struct {
	MovieService service.MovieService
}

func NewMovieHandler(movieService service.MovieService) MovieHandler {
	return &MovieHandlerImpl{
		MovieService: movieService,
	}
}

func (m MovieHandlerImpl) Create(writer http.ResponseWriter, request *http.Request) {
	movieCreateRequest := web.MovieRequest{}
	helper.ReadFromRequestBody(request, &movieCreateRequest)
	movieResponse := m.MovieService.Create(movieCreateRequest)

	webResponse := web.WebResponse{
		Code:   http.StatusCreated,
		Error:  nil,
		Result: movieResponse,
	}
	helper.WriteToResponseBody(writer, http.StatusCreated, webResponse)
}

func (m MovieHandlerImpl) Update(writer http.ResponseWriter, request *http.Request, params chi.RouteParams) {
	panic("implement me")
}

func (m MovieHandlerImpl) Delete(writer http.ResponseWriter, request *http.Request, params chi.RouteParams) {
	panic("implement me")
}

func (m MovieHandlerImpl) FindBySlug(writer http.ResponseWriter, request *http.Request) {
	paramSlug := chi.URLParam(request, "slug")

	movieResponse := m.MovieService.FindBySlug(paramSlug)

	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Error:  nil,
		Result: movieResponse,
	}
	helper.WriteToResponseBody(writer, http.StatusOK, webResponse)
}
