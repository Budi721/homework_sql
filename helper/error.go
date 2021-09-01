package helper

import (
	"errors"
	"github.com/Budi721/homework_sql/model/web"
	my "github.com/go-mysql/errors"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
	"net/http"
)

func HandleBadRequestRequest(w http.ResponseWriter, err error) bool {
	exception, ok := err.(validator.ValidationErrors)
	if ok {
		webResponse := web.WebResponse{
			Code:   http.StatusBadRequest,
			Error:  exception.Error(),
			Result: nil,
		}

		WriteToResponseBody(w, http.StatusBadRequest, webResponse)
		return true
	}
	return false
}

func HandleUnprocessableEntity(w http.ResponseWriter, err error) bool {
	if ok, myerr := my.Error(err); ok { // MySQL error
		if myerr == my.ErrDupeKey {
			webResponse := web.WebResponse{
				Code:   http.StatusUnprocessableEntity,
				Error:  my.ErrDupeKey.Error(),
				Result: nil,
			}

			WriteToResponseBody(w, http.StatusUnprocessableEntity, webResponse)
			return true
		}
	}
	return false
}

func HandleNotFound(w http.ResponseWriter, err error) bool {
	if errors.Is(err, gorm.ErrRecordNotFound) {
		webResponse := web.WebResponse{
			Code:   http.StatusNotFound,
			Error:  err.Error(),
			Result: nil,
		}

		WriteToResponseBody(w, http.StatusNotFound, webResponse)
		return true
	}
	return false
}