package main

import (
	"log"
	"net/http"
	"time"

	"github.com/Ashis-007/hms/config"
	"github.com/Ashis-007/hms/db"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/jmoiron/sqlx"
)

type AppConfig struct {
	db *sqlx.DB
}

var App AppConfig

func main() {
	// load config
	env, err := config.LoadEnv("../../")
	if err != nil {
		return
	}

	App = AppConfig{}

	err = runServer(env, &App)
	if err != nil {
		log.Fatal("an error occurred while running server: ", err)
		return
	}
}

func runServer(env config.Env, app *AppConfig) error {
	router := createRouter()

	appDB, err := db.ConnectDB(env)
	if err != nil {
		return err
	}
	defer appDB.Close()

	app.db = appDB

	// start server
	http.ListenAndServe(env.Port, router)

	return nil
}

func createV1Router() chi.Router {
	v1Router := chi.NewRouter()
	return v1Router
}

func createRouter() chi.Router {
	// initialize chi router
	router := chi.NewRouter()

	// middlewares
	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PATCH", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	// Set a timeout value on the request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped.
	router.Use(middleware.Timeout(60 * time.Second))

	v1Router := createV1Router()

	router.Mount("/v1", v1Router)

	return router
}
