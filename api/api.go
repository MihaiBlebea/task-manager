package api

import (
	"net/http"
)

type Logger interface {
	Info(args ...interface{})
	Trace(args ...interface{})
	Debug(args ...interface{})
	Print(args ...interface{})
	Warn(args ...interface{})
	Warning(args ...interface{})
	Error(args ...interface{})
	Fatal(args ...interface{})
	Panic(args ...interface{})
}

type Handler interface {
	HealthEndpoint() http.Handler

	SelectProjectEndpoint() http.Handler
	SelectUserProjectsEndpoint() http.Handler
	CreateProjectEndpoint() http.Handler
	DeleteProjectEndpoint() http.Handler
	UpdateProjectEndpoint() http.Handler

	CreateTaskEndpoint() http.Handler
	DeleteTaskEndpoint() http.Handler
	CompleteTaskEndpoint() http.Handler
}

type Service struct {
	logger  Logger
	handler Handler
}

func New(handler Handler, logger Logger) *Service {
	return &Service{logger, handler}
}

func (s *Service) Server() {
	server(s.handler, s.logger)
}
