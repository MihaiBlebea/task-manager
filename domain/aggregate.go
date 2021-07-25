package domain

import (
	"errors"

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
	RegisterUser(username, firstName, lastName string, chatID int) error
	IsRegistered(chatID int) bool
	DeleteUser(chatID int) error
	AllUser() ([]user.User, error)
}

type taskManager struct {
	userRepo user.Repo
}

func New(conn *gorm.DB) TaskManager {
	return &taskManager{
		userRepo: user.NewRepo(conn),
	}
}

func (tm *taskManager) RegisterUser(username, firstName, lastName string, chatID int) error {
	user := user.New(username, firstName, lastName, chatID)

	_, err := tm.userRepo.Save(user)
	if err != nil {
		return err
	}

	return nil
}

func (tm *taskManager) IsRegistered(chatID int) bool {
	_, err := tm.userRepo.FindWithChatID(chatID)
	if err != nil {
		return false
	}

	return true
}

func (tm *taskManager) DeleteUser(chatID int) error {
	err := tm.userRepo.DeleteByChatID(chatID)
	if err != nil {
		return err
	}

	return nil
}

func (tm *taskManager) AllUser() ([]user.User, error) {
	return tm.userRepo.All()
}
