package main

import (
	"log"
	"schoolsystem/learninghub-api/core"
	"schoolsystem/learninghub-api/db"
	"schoolsystem/learninghub-api/middlewares"

	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/casbin/casbin"

	"schoolsystem/learninghub-api/course"
	courseRepository "schoolsystem/learninghub-api/course/repository"
	courseUsecase "schoolsystem/learninghub-api/course/usecase"

	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/rs/cors"
)

var (
	s3Server *s3.S3
)

func main() {
	// setup logger
	logger, err := core.NewLogger()
	if err != nil {
		logger.Panic(err.Error())
	}

	// setup config
	conf, err := core.NewConfig()
	if err != nil {
		logger.Panic(err.Error())
	}
	log.Println(conf.String())

	// database connection
	dbConn, err := db.NewDatabaseConnection(conf.Database.URL)

	if err != nil {
		logger.Panic(err.Error())
	}

	log.Println("Successfully connected to mlab.")

	defer dbConn.Close()

	// setting up casbin authorization
	e, err := casbin.NewEnforcerSafe(conf.CasbinConfPath+"model.conf", conf.CasbinConfPath+"policy.csv")

	if err != nil {
		logger.Panic(err.Error())
	}

	// setup routes and middleware
	r := chi.NewRouter()

	r.Use(middleware.StripSlashes)
	r.Use(middlewares.Authorizer(e))

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.SetHeader("Content-Type", "application/json"))

	courseRepo := courseRepository.NewMongoCourseRepository(dbConn)
	courseUsecase := courseUsecase.NewCourseUsecase(courseRepo)
	course.NewCourseHttpHandler(r, courseUsecase)

	r.HandleFunc("/p", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("learninghub-api"))
	})

	handler := cors.Default().Handler(r)
	log.Printf("Server running on: %s", conf.BaseURL)
	http.ListenAndServe(conf.BaseURL, handler)

}
