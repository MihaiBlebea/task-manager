package api

import (
	"fmt"
	"log"

	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

const prefix = "/api/v1"

func server(handler Handler, logger Logger) {
	r := mux.NewRouter()

	api := r.PathPrefix(prefix).Subrouter()

	// Handle api calls
	api.Handle("/health-check", handler.HealthEndpoint()).
		Methods(http.MethodGet)

	// User
	api.Handle("/user/register", handler.RegisterEndpoint()).
		Methods(http.MethodPost)

	api.Handle("/user/login", handler.LoginEndpoint()).
		Methods(http.MethodPost)

	// Project
	api.Handle("/project", handler.CreateProjectEndpoint()).
		Methods(http.MethodPost)

	api.Handle("/project/{project_id}", handler.SelectProjectEndpoint()).
		Methods(http.MethodGet)

	api.Handle("/projects/user/{user_id}", handler.SelectUserProjectsEndpoint()).
		Methods(http.MethodGet)

	api.Handle("/project/{project_id}", handler.DeleteProjectEndpoint()).
		Methods(http.MethodDelete)

	api.Handle("/project/{project_id}", handler.UpdateProjectEndpoint()).
		Methods(http.MethodPut)

	// Task
	api.Handle("/task", handler.CreateTaskEndpoint()).
		Methods(http.MethodPost)

	api.Handle("/task/{task_id}", handler.UpdateTaskEndpoint()).
		Methods(http.MethodPut)

	api.Handle("/task/{task_id}", handler.DeleteTaskEndpoint()).
		Methods(http.MethodDelete)

	api.Handle("/task/complete/{task_id}", handler.CompleteTaskEndpoint()).
		Methods(http.MethodPut)

	r.Use(loggerMiddleware(logger))

	srv := &http.Server{
		Handler:      cors.AllowAll().Handler(api),
		Addr:         fmt.Sprintf("0.0.0.0:%s", os.Getenv("HTTP_PORT")),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	logger.Info(fmt.Sprintf("Started server on port %s", os.Getenv("HTTP_PORT")))

	log.Fatal(srv.ListenAndServe())
}
