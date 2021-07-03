package handler

import (
	"net/http"

	"github.com/MihaiBlebea/task-manager/api/handlers/project"
	"github.com/MihaiBlebea/task-manager/domain"
	"gorm.io/gorm"
)

type Service struct {
	domain domain.TaskManager
}

func New(conn *gorm.DB) *Service {
	return &Service{
		domain: domain.New(conn),
	}
}

func (s *Service) HealthEndpoint() http.Handler {
	return healthEndpoint()
}

func (s *Service) SelectProjectEndpoint() http.Handler {
	return project.SelectIdHandler(s.domain)
}

func (s *Service) SelectUserProjectsEndpoint() http.Handler {
	return project.SelectUserHandler(s.domain)
}

func (s *Service) CreateProjectEndpoint() http.Handler {
	return project.CreateHandler(s.domain)
}

func (s *Service) DeleteProjectEndpoint() http.Handler {
	return project.DeleteHandler(s.domain)
}

func (s *Service) UpdateProjectEndpoint() http.Handler {
	return project.UpdateHandler(s.domain)
}
