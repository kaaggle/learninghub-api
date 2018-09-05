package main

import (
	"schoolsystem/learninghub-api/core"
	"schoolsystem/learninghub-api/db"

	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {
	// setup logger
	logger, err := core.NewLogger()
	if err != nil {
		logger.Panic(err.Error())
	}

	// setup config
	conf := core.NewConfig()

	// database connection
	dbConn, err := db.NewDatabaseConnection(conf.Database.URL)

	if err != nil {
		logger.Panic(err.Error())
	}

	defer dbConn.Close()

	// setup routes and middleware
	r := chi.NewRouter()

	r.Use(middleware.StripSlashes)
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.SetHeader("Content-Type", "application/json"))

	r.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("schoolview api"))
	})

	http.ListenAndServe(conf.BaseURL, r)
}
