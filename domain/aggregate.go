package domain

import (
	"github.com/MihaiBlebea/task-manager/domain/project"
	"gorm.io/gorm"
)

type TaskManager interface {
	GetProject(userID, projectID int) (*project.Project, error)
	GetUserProjects(userID int) ([]project.Project, error)
	CreateNewProject(userID int, title, color, description, icon string) (int, error)
	DeleteProject(userID, projectID int) error
	UpdateProject(userID, projectID int, title, color, description, icon string) error
}

type taskManager struct {
	projectRepo project.Repo
}

func New(conn *gorm.DB) TaskManager {
	return &taskManager{
		projectRepo: project.NewRepo(conn),
	}
}

func (tm *taskManager) GetProject(userID, projectID int) (*project.Project, error) {
	return tm.projectRepo.FindWithID(projectID)
}

func (tm *taskManager) GetUserProjects(userID int) ([]project.Project, error) {
	return tm.projectRepo.FindWithUserID(userID)
}

func (tm *taskManager) CreateNewProject(userID int, title, color, description, icon string) (int, error) {
	proj := project.New(userID, title)
	err := tm.projectRepo.Save(proj)

	return proj.ID, err
}

func (tm *taskManager) DeleteProject(userID, projectID int) error {
	return tm.projectRepo.Delete(projectID)
}

func (tm *taskManager) UpdateProject(userID, projectID int, title, color, description, icon string) error {
	proj := project.Project{
		ID:          projectID,
		UserID:      userID,
		Title:       title,
		Color:       color,
		Description: description,
		Icon:        icon,
	}

	return tm.projectRepo.Update(&proj)
}
