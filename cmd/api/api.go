package main

import (
	"log"
	"net/http"
	"time"

	httpHandlers "social/internal/http"
	"social/internal/store"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type application struct {
	config config
	store  store.Storage
}

type config struct {
	addr string
	db   dbConfig
}

type dbConfig struct {
	addr               string
	maxOpenConnections int
	maxIdleConnections int
	maxIdleTime        string
}

func (app *application) mount() *chi.Mux {
	r := chi.NewRouter()

	// A good base middleware stack
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Set a timeout value on the request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped.
	r.Use(middleware.Timeout(60 * time.Second))

	r.Route("/api/v1", func(r chi.Router) {
		r.Get("/health", app.healthCheckHandler)

		// mounting endpoints

		// first we create a user handler instance
		userHandler := httpHandlers.UserHandler{
			Store: app.store.Users,
		}

		postHandler := httpHandlers.PostHandler{
			Store: app.store.Posts,
		}

		r.Mount("/users", userHandler.UserRoutes())
		r.Mount("/posts", postHandler.PostRoutes())
	})

	// Use chi.Walk to print the routes
	// chi.Walk(r, func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
	// 	fmt.Printf("[%s] %s\n", method, route)
	// 	return nil
	// })

	return r
}

func (app *application) run(mux *chi.Mux) error {
	srv := &http.Server{
		Addr:    app.config.addr,
		Handler: mux,
		// if our server takes more than 30 seconds to write a response to the client, we time it out
		WriteTimeout: time.Second * 30,

		// if the client takes more than 10 seconds to read the server's response, we time it out
		ReadTimeout: time.Second * 10,
		IdleTimeout: time.Minute,
	}

	log.Printf("Server has started at %s", app.config.addr)
	return srv.ListenAndServe()
}
