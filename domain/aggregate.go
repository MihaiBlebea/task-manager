package domain

import (
	"time"

	"github.com/MihaiBlebea/task-manager/domain/project"
	"github.com/MihaiBlebea/task-manager/domain/task"
	"gorm.io/gorm"
)

const (
	timeLayout = "2006-01-02T15:04:05.000Z"
)

type TaskManager interface {
	GetProject(userID, projectID int) (*project.Project, error)
	GetUserProjects(userID int) ([]project.Project, error)
	CreateProject(userID int, title, color, description, icon string) (int, error)
	DeleteProject(userID, projectID int) error
	UpdateProject(userID, projectID int, title, color, description, icon string) error
	CreateTask(
		userID,
		subtaskID,
		projectID int,
		title,
		note string,
		expire string,
		repeat bool,
		reapeatDayOfWeek int,
		repeatTimeOfDay string,
		priority int) (int, error)
	DeleteTask(userID, taskID int) error
	CompleteTask(userID, taskID int) error
}

type taskManager struct {
	projectRepo project.Repo
	taskRepo    task.Repo
}

func New(conn *gorm.DB) TaskManager {
	return &taskManager{
		projectRepo: project.NewRepo(conn),
		taskRepo:    task.NewRepo(conn),
	}
}

func (tm *taskManager) GetProject(userID, projectID int) (*project.Project, error) {
	project, err := tm.projectRepo.FindWithID(projectID)
	if err != nil {
		return project, err
	}

	// Fetch all the tasks associated with this project
	tasks, err := tm.taskRepo.FindTasksWithProjectID(project.ID)
	if err != nil {
		return project, err
	}

	if len(tasks) == 0 {
		return project, nil
	}

	project.Tasks = tasks

	return project, nil
}

func (tm *taskManager) GetUserProjects(userID int) ([]project.Project, error) {
	return tm.projectRepo.FindWithUserID(userID)
}

func (tm *taskManager) CreateProject(
	userID int,
	title,
	color,
	description,
	icon string) (int, error) {
	proj := project.New(userID, title)

	return tm.projectRepo.Save(proj)
}

func (tm *taskManager) DeleteProject(userID, projectID int) error {
	return tm.projectRepo.Delete(projectID)
}

func (tm *taskManager) UpdateProject(
	userID,
	projectID int,
	title,
	color,
	description,
	icon string) error {
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

func (tm *taskManager) CreateTask(
	userID,
	subtaskID,
	projectID int,
	title,
	note string,
	expire string,
	repeat bool,
	reapeatDayOfWeek int,
	repeatTimeOfDay string,
	priority int) (int, error) {

	expireTime, err := time.Parse(timeLayout, expire)
	if err != nil {
		return 0, err
	}

	return tm.taskRepo.Save(&task.Task{
		SubtaskID:       subtaskID,
		ProjectID:       projectID,
		Title:           title,
		Note:            note,
		Expire:          expireTime,
		Repeat:          repeat,
		RepeatDayOfWeek: reapeatDayOfWeek,
		RepeatTimeOfDay: repeatTimeOfDay,
		Priority:        priority,
	})
}

func (tm *taskManager) DeleteTask(userID, taskID int) error {
	return tm.taskRepo.Delete(taskID)
}

func (tm *taskManager) CompleteTask(userID, taskID int) error {
	task, err := tm.taskRepo.FindWithID(taskID)
	if err != nil {
		return err
	}
	task.Completed = true

	return tm.taskRepo.Update(task)
}
