package main

import (
	"github.com/Budi721/homework_sql/app"
	"github.com/Budi721/homework_sql/config"
	"github.com/Budi721/homework_sql/handler"
	"github.com/Budi721/homework_sql/repository"
	"github.com/Budi721/homework_sql/router"
	"github.com/Budi721/homework_sql/service"
	"github.com/go-playground/validator"
	"log"
	"net/http"
	"strconv"
)

func main() {
	cfg := config.Init()

	db, _ := app.InitDB(); app.Migrate(db)
	log.Printf("Starting up on http://localhost:%v", cfg.AppPort)

	validate := validator.New()
	movieRepository := repository.NewMovieRepository()
	movieService := service.NewMovieService(movieRepository, db, validate)
	movieHandler := handler.NewMovieHandler(movieService)
	newRouter := router.NewRouter(movieHandler)

	log.Fatal(http.ListenAndServe(":" + strconv.Itoa(cfg.AppPort), newRouter))
}
