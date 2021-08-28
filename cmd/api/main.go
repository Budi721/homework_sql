package main

import (
	"github.com/Budi721/homework_sql/app"
	"github.com/Budi721/homework_sql/config"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"log"
	"net/http"
	"strconv"
)

func main() {
	cfg := config.Init()

	db, _ := app.InitDB(); app.Migrate(db)
	log.Printf("Starting up on http://localhost:%v", cfg.AppPort)

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		w.Write([]byte("Hello World!"))
	})

	log.Fatal(http.ListenAndServe(":" + strconv.Itoa(cfg.AppPort), r))
}
