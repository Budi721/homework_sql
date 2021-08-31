package main

import (
	"github.com/Budi721/homework_sql/app"
	"github.com/Budi721/homework_sql/handler"
	"github.com/Budi721/homework_sql/repository"
	"github.com/Budi721/homework_sql/router"
	"github.com/Budi721/homework_sql/service"
	"github.com/go-playground/validator/v10"
	"log"
	"net/http"
	"strconv"
)

func main() {
	init := app.Init()
	log.Printf("Starting up on http://localhost:%v", init.Config.AppPort)

	validate := validator.New()
	movieRepository := repository.NewMovieRepository()
	movieService := service.NewMovieService(movieRepository, init.DbConn, validate)
	movieHandler := handler.NewMovieHandler(movieService)
	newRouter := router.NewRouter(movieHandler)

	log.Fatal(http.ListenAndServe(":" + strconv.Itoa(init.Config.AppPort), newRouter))
}
