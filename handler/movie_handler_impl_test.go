package handler_test

import (
	"encoding/json"
	"github.com/Budi721/homework_sql/app"
	"github.com/Budi721/homework_sql/handler"
	"github.com/Budi721/homework_sql/model/domain"
	"github.com/Budi721/homework_sql/repository"
	"github.com/Budi721/homework_sql/router"
	"github.com/Budi721/homework_sql/service"
	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func setupRouter(db *gorm.DB) http.Handler {
	validate := validator.New()
	movieRepository := repository.NewMovieRepository()
	movieService := service.NewMovieService(movieRepository, db, validate)
	movieHandler := handler.NewMovieHandler(movieService)

	return router.NewRouter(movieHandler)
}

func truncateMovie(db *gorm.DB) {
	db.Exec("TRUNCATE movie")
}

func TestCreateCategorySuccess(t *testing.T) {
	db, _ := app.InitDB()
	truncateMovie(db)
	routes := setupRouter(db)

	requestBody := strings.NewReader(`{ "title": "Titanic23","slug": "jh","description": "lorem ipsum","duration": 10,"image": "titanic poster url"}`)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:8000/movies", requestBody)
	request.Header.Add("Content-Type", "application/json")

	recorder := httptest.NewRecorder()

	routes.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 201, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 201, int(responseBody["code"].(float64)))
}

func TestGetCategory(t *testing.T) {
	db, _ := app.InitDB()
	truncateMovie(db)
	routes := setupRouter(db)

	request := httptest.NewRequest(http.MethodGet, "http://localhost:8000/movies/jh", nil)

	recorder := httptest.NewRecorder()

	routes.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	assert.Equal(t, nil, responseBody["error"])
}

func TestDeleteCategorySuccess(t *testing.T) {
	db, _ := app.InitDB()
	truncateMovie(db)

	db.Create(domain.Movie{
		Title:       "test",
		Slug:        "slug",
		Description: "sdahjjdsa",
		Duration:    8,
		Image:       "sdakjhsa",
	})

	router := setupRouter(db)

	request := httptest.NewRequest(http.MethodDelete, "http://localhost:3000/movies/slug", nil)
	request.Header.Add("Content-Type", "application/json")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 200, int(responseBody["code"].(float64)))
}
