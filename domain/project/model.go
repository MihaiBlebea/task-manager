package project

import (
	"time"

	"github.com/MihaiBlebea/task-manager/domain/task"
	"gorm.io/gorm"
)

type Project struct {
	ID          int         `json:"id"`
	UserID      int         `json:"user_id"`
	Title       string      `json:"title"`
	Color       string      `json:"color"`
	Description string      `json:"description"`
	Icon        string      `json:"icon"`
	Tasks       []task.Task `json:"tasks" gorm:"-"`
	CreatedAt   time.Time   `json:"created_at"`
	UpdatedAt   time.Time   `json:"updated_at"`
}

func New(userID int, title string) *Project {
	return &Project{UserID: userID, Title: title}
}

func (p *Project) AfterDelete(conn *gorm.DB) (err error) {
	return conn.Where("project_id = ?", p.ID).Delete(&task.Task{}).Error
}
