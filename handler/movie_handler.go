package handler

import (
	"net/http"
)

type MovieHandler interface {
	Create(writer http.ResponseWriter, request *http.Request)
	Update(writer http.ResponseWriter, request *http.Request)
	Delete(writer http.ResponseWriter, request *http.Request)
	FindBySlug(writer http.ResponseWriter, request *http.Request)
}
