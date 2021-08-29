package handler

import (
	"github.com/go-chi/chi"
	"net/http"
)

type MovieHandler interface {
	Create(writer http.ResponseWriter, request *http.Request)
	Update(writer http.ResponseWriter, request *http.Request, params chi.RouteParams)
	Delete(writer http.ResponseWriter, request *http.Request, params chi.RouteParams)
	FindById(writer http.ResponseWriter, request *http.Request, params chi.RouteParams)
}
