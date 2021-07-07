package domain

import (
	"errors"
	"time"

	"github.com/MihaiBlebea/task-manager/domain/project"
	"github.com/MihaiBlebea/task-manager/domain/task"
	"github.com/MihaiBlebea/task-manager/domain/user"
	"gorm.io/gorm"
)

const (
	timeLayout = "2006-01-02T15:04:05.000Z"
)

var (
	ErrInvalidUserID = errors.New("Invalid user ID")
	ErrUserNotOwner  = errors.New("User is not the owner")
)

type TaskManager interface {
	RegisterUser(username, email, password string) (int, string, error)
	LoginUser(email, password string) (int, string, error)
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
	userRepo    user.Repo
	projectRepo project.Repo
	taskRepo    task.Repo
}

func New(conn *gorm.DB) TaskManager {
	return &taskManager{
		userRepo:    user.NewRepo(conn),
		projectRepo: project.NewRepo(conn),
		taskRepo:    task.NewRepo(conn),
	}
}

func (tm *taskManager) RegisterUser(username, email, password string) (int, string, error) {
	user, err := user.New(username, email, password)
	if err != nil {
		return 0, "", err
	}

	id, err := tm.userRepo.Save(user)
	if err != nil {
		return 0, "", err
	}

	_, err = user.GenerateJWT()
	if err != nil {
		return 0, "", err
	}

	if err := tm.userRepo.Update(user); err != nil {
		return 0, "", err
	}

	return id, user.Token, nil
}

func (tm *taskManager) LoginUser(email, password string) (int, string, error) {
	u, err := tm.userRepo.FindWithEmail(email)
	if err != nil {
		return 0, "", err
	}

	if u.CheckPasswordHash(password) == false {
		return 0, "", errors.New("Could not auth user")
	}

	return u.ID, u.Token, nil
}

func (tm *taskManager) GetProject(userID, projectID int) (*project.Project, error) {
	if isValid := tm.validateUserID(userID); isValid == false {
		return &project.Project{}, ErrInvalidUserID
	}

	proj, err := tm.projectRepo.FindWithID(projectID)
	if err != nil {
		return &project.Project{}, err
	}

	// Validate if the user is the owner of the project
	if proj.ID != userID {
		return &project.Project{}, ErrUserNotOwner
	}

	// Fetch all the tasks associated with this project
	tasks, err := tm.taskRepo.FindTasksWithProjectID(proj.ID)
	if err != nil {
		return proj, err
	}

	if len(tasks) == 0 {
		return proj, nil
	}

	proj.Tasks = tasks

	return proj, nil
}

func (tm *taskManager) GetUserProjects(userID int) (projects []project.Project, _ error) {
	if isValid := tm.validateUserID(userID); isValid == false {
		return projects, ErrInvalidUserID
	}

	return tm.projectRepo.FindWithUserID(userID)
}

func (tm *taskManager) CreateProject(
	userID int,
	title,
	color,
	description,
	icon string) (int, error) {
	if isValid := tm.validateUserID(userID); isValid == false {
		return 0, ErrInvalidUserID
	}

	proj := project.New(userID, title)
	proj.Color = color
	proj.Description = description
	proj.Icon = icon

	return tm.projectRepo.Save(proj)
}

func (tm *taskManager) DeleteProject(userID, projectID int) error {
	if isValid := tm.validateUserID(userID); isValid == false {
		return ErrInvalidUserID
	}

	if isOwner := tm.validateProjectOwner(userID, projectID); isOwner == false {
		return ErrUserNotOwner
	}

	return tm.projectRepo.Delete(projectID)
}

func (tm *taskManager) UpdateProject(
	userID,
	projectID int,
	title,
	color,
	description,
	icon string) error {

	if isValid := tm.validateUserID(userID); isValid == false {
		return ErrInvalidUserID
	}

	if isOwner := tm.validateProjectOwner(userID, projectID); isOwner == false {
		return ErrUserNotOwner
	}

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
