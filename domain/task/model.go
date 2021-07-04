package task

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

type Task struct {
	ID              int       `json:"id"`
	SubtaskID       int       `json:"subtask_id"`
	ProjectID       int       `json:"project_id"`
	Title           string    `json:"title"`
	Note            string    `json:"note"`
	Expire          time.Time `json:"expire"`
	Repeat          bool      `json:"repeat"`
	RepeatDayOfWeek int       `json:"repeat_day_of_week"`
	RepeatTimeOfDay string    `json:"repeat_time_of_day"`
	Priority        int       `json:"priority"`
	Completed       bool      `json:"completed"`
	Subtasks        []Task    `json:"subtasks" gorm:"-"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

func New(projectID int, title string) *Task {
	return &Task{ProjectID: projectID, Title: title}
}

func (t *Task) AfterFind(conn *gorm.DB) (err error) {
	if t.SubtaskID == 0 {
		subtasks := make([]Task, 0)
		err := conn.Where("subtask_id = ?", t.ID).Find(&subtasks).Error
		if err != nil {
			return err
		}

		t.Subtasks = subtasks
	}

	return
}

func (t *Task) BeforeCreate(conn *gorm.DB) (err error) {
	// This is a main task, no need for other validation
	if t.SubtaskID == 0 {
		return
	}

	// This is a subtask
	// Check if there is a main task to associate with this subtask
	task := &Task{}
	if err := conn.Where("id = ?", t.SubtaskID).Find(&task).Error; err != nil {
		return err
	}

	if task.ID == 0 {
		return ErrNoRecord
	}

	// Update the main task to have completed = false
	if err := conn.Model(&Task{}).Where("id = ? AND completed = true", task.ID).Update("completed", false).Error; err != nil {
		return err
	}

	return
}

func (t *Task) AfterDelete(conn *gorm.DB) (err error) {
	if t.SubtaskID == 0 && len(t.Subtasks) > 0 {
		return conn.Delete(&t.Subtasks).Error
	}

	return
}

func (t *Task) AfterUpdate(conn *gorm.DB) (err error) {
	// This is the main task
	if t.SubtaskID == 0 {
		if len(t.Subtasks) > 0 && t.Completed == true {
			if err := conn.Model(&Task{}).Where("subtask_id = ? AND completed = false", t.ID).Update("completed", true).Error; err != nil {
				return err
			}
		}

		return
	}

	// This is a subtask and has been updated
	subtasks := make([]Task, 0)
	if err := conn.Where("subtask_id = ?", t.SubtaskID).Find(&subtasks).Error; err != nil {
		return err
	}

	if len(subtasks) > 0 {
		completed := true
		for _, subtask := range subtasks {
			if subtask.Completed == false {
				completed = false
			}
		}

		if completed == false {
			fmt.Println("Not all tasks are completed")
			return
		}

		if err := conn.Model(&Task{}).Where("subtask_id = ? OR id = ? AND completed = false", t.SubtaskID, t.SubtaskID).Update("completed", true).Error; err != nil {
			return err
		}
	}

	return
}
