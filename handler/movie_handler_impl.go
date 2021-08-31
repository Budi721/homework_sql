package handler

import (
	"errors"
	"github.com/Budi721/homework_sql/helper"
	"github.com/Budi721/homework_sql/model/web"
	"github.com/Budi721/homework_sql/service"
	"github.com/go-chi/chi"
	my "github.com/go-mysql/errors"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
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
	exception, ok := err.(validator.ValidationErrors)
	if ok {
		webResponse := web.WebResponse{
			Code:   http.StatusBadRequest,
			Error:  exception.Error(),
			Result: nil,
		}

		helper.WriteToResponseBody(writer, http.StatusBadRequest, webResponse)
		return
	}

	if ok, myerr := my.Error(err); ok { // MySQL error
		if myerr == my.ErrDupeKey {
			webResponse := web.WebResponse{
				Code:   http.StatusUnprocessableEntity,
				Error:  my.ErrDupeKey.Error(),
				Result: nil,
			}

			helper.WriteToResponseBody(writer, http.StatusUnprocessableEntity, webResponse)
			return
		}
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
	exception, ok := err.(validator.ValidationErrors)
	if ok {
		webResponse := web.WebResponse{
			Code:   http.StatusBadRequest,
			Error:  exception.Error(),
			Result: nil,
		}

		helper.WriteToResponseBody(writer, http.StatusBadRequest, webResponse)
		return
	}

	if ok, myerr := my.Error(err); ok { // MySQL error
		if myerr == my.ErrDupeKey {
			webResponse := web.WebResponse{
				Code:   http.StatusUnprocessableEntity,
				Error:  my.ErrDupeKey.Error(),
				Result: nil,
			}

			helper.WriteToResponseBody(writer, http.StatusUnprocessableEntity, webResponse)
			return
		}
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
	if errors.Is(err, gorm.ErrRecordNotFound) {
		webResponse := web.WebResponse{
			Code:   http.StatusNotFound,
			Error:  err.Error(),
			Result: nil,
		}

		helper.WriteToResponseBody(writer, http.StatusNotFound, webResponse)
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
	if errors.Is(err, gorm.ErrRecordNotFound) {
		webResponse := web.WebResponse{
			Code:   http.StatusNotFound,
			Error:  err.Error(),
			Result: nil,
		}

		helper.WriteToResponseBody(writer, http.StatusNotFound, webResponse)
		return
	}

	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Error:  nil,
		Result: movieResponse,
	}
	helper.WriteToResponseBody(writer, http.StatusOK, webResponse)
}
