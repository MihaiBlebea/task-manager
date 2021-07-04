package task

import (
	"errors"

	"gorm.io/gorm"
)

var (
	ErrNoRecord    error = errors.New("Record not found")
	ErrNoRecords   error = errors.New("Records not found")
	ErrNotInserted error = errors.New("Could not insert record")
)

type repo struct {
	conn *gorm.DB
}

type Repo interface {
	FindWithID(taskID int) (*Task, error)
	FindWithProjectID(projectID int) ([]Task, error)
	FindTasksWithProjectID(projectID int) ([]Task, error)
	FindSubtasksWithID(subtaskID int) ([]Task, error)
	Save(task *Task) (int, error)
	Delete(taskID int) error
	DeleteTasks(taskIDs []int) error
	Update(task *Task) error
}

func NewRepo(conn *gorm.DB) Repo {
	return &repo{conn}
}

func (r *repo) FindWithID(taskID int) (*Task, error) {
	task := Task{}
	err := r.conn.Where("id = ?", taskID).Find(&task).Error
	if err != nil {
		return &task, err
	}

	if task.ID == 0 {
		return &task, ErrNoRecord
	}

	return &task, nil
}

// Find all with project id, tasks and subtasks
func (r *repo) FindWithProjectID(projectID int) ([]Task, error) {
	tasks := make([]Task, 0)
	err := r.conn.Where("project_id = ?", projectID).Find(&tasks).Error
	if err != nil {
		return tasks, err
	}

	return tasks, nil
}

// Find just tasks with project id, no subtasks
func (r *repo) FindTasksWithProjectID(projectID int) ([]Task, error) {
	tasks := make([]Task, 0)
	err := r.conn.Where("project_id = ? AND subtask_id = 0", projectID).Find(&tasks).Error
	if err != nil {
		return tasks, err
	}

	return tasks, nil
}

func (r *repo) FindSubtasksWithID(subtaskID int) ([]Task, error) {
	tasks := make([]Task, 0)
	err := r.conn.Where("subtask_id = ?", subtaskID).Find(&tasks).Error
	if err != nil {
		return tasks, err
	}

	return tasks, nil
}

func (r *repo) Save(task *Task) (int, error) {
	cmd := r.conn.Create(task)
	if cmd.RowsAffected == 0 {
		return 0, ErrNotInserted
	}

	return task.ID, cmd.Error
}

func (r *repo) Delete(taskID int) error {
	task := &Task{}
	if err := r.conn.Where("id = ?", taskID).Find(task).Error; err != nil {
		return err
	}

	cmd := r.conn.Delete(task)
	if cmd.RowsAffected == 0 {
		return ErrNoRecord
	}

	return cmd.Error
}

func (r *repo) DeleteTasks(taskIDs []int) error {
	tasks := make([]Task, 0)
	if err := r.conn.Where("id IN ?", taskIDs).Find(&tasks).Error; err != nil {
		return err
	}

	cmd := r.conn.Delete(&tasks)
	if cmd.RowsAffected == 0 {
		return ErrNoRecord
	}

	return cmd.Error
}

func (r *repo) Update(task *Task) error {
	cmd := r.conn.Model(task).Updates(task)
	if cmd.RowsAffected == 0 {
		return ErrNoRecord
	}

	return cmd.Error
}
