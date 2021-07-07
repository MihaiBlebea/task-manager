package handler

import (
	"net/http"

	"github.com/MihaiBlebea/task-manager/api/handlers/project"
	"github.com/MihaiBlebea/task-manager/api/handlers/task"
	"github.com/MihaiBlebea/task-manager/api/handlers/user"
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

func (s *Service) RegisterEndpoint() http.Handler {
	return user.RegisterHandler(s.domain)
}

func (s *Service) LoginEndpoint() http.Handler {
	return user.LoginHandler(s.domain)
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

func (s *Service) CreateTaskEndpoint() http.Handler {
	return task.CreateHandler(s.domain)
}

func (s *Service) DeleteTaskEndpoint() http.Handler {
	return task.DeleteHandler(s.domain)
}

func (s *Service) CompleteTaskEndpoint() http.Handler {
	return task.CompleteHandler(s.domain)
}
