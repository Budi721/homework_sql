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
	movieResponse, err := m.MovieService.Create(movieCreateRequest)
	if helper.HandleBadRequestRequest(writer, err) {
		return
	}
	if helper.HandleUnprocessableEntity(writer, err) {
		return
	}

	webResponse := web.WebResponse{
		Code:   http.StatusCreated,
		Error:  nil,
		Result: movieResponse,
	}
	helper.WriteToResponseBody(writer, http.StatusCreated, webResponse)
}

func (m MovieHandlerImpl) Update(writer http.ResponseWriter, request *http.Request) {
	paramSlug := chi.URLParam(request, "slug")
	movieUpdateRequest := web.MovieRequest{}
	helper.ReadFromRequestBody(request, &movieUpdateRequest)
	movieResponse, err := m.MovieService.Update(movieUpdateRequest, paramSlug)
	if helper.HandleBadRequestRequest(writer, err) {
		return
	}
	if helper.HandleUnprocessableEntity(writer, err) {
		return
	}

	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Error:  nil,
		Result: movieResponse,
	}
	helper.WriteToResponseBody(writer, http.StatusOK, webResponse)
}

func (m MovieHandlerImpl) Delete(writer http.ResponseWriter, request *http.Request) {
	paramSlug := chi.URLParam(request, "slug")

	err := m.MovieService.Delete(paramSlug)
	if helper.HandleNotFound(writer, err) {
		return
	}
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Error:  nil,
		Result: "Success",
	}
	helper.WriteToResponseBody(writer, http.StatusOK, webResponse)
}

func (m MovieHandlerImpl) FindBySlug(writer http.ResponseWriter, request *http.Request) {
	paramSlug := chi.URLParam(request, "slug")

	movieResponse, err := m.MovieService.FindBySlug(paramSlug)
	if helper.HandleNotFound(writer, err) {
		return
	}

	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Error:  nil,
		Result: movieResponse,
	}
	helper.WriteToResponseBody(writer, http.StatusOK, webResponse)
}
