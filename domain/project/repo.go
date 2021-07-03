package project

import (
	"errors"

	"gorm.io/gorm"
)

var (
	ErrNoRecord  error = errors.New("Record not found")
	ErrNoRecords error = errors.New("Records not found")
)

type repo struct {
	conn *gorm.DB
}

type Repo interface {
	FindWithID(projectID int) (*Project, error)
	FindWithUserID(userID int) ([]Project, error)
	Save(project *Project) error
	Delete(projectID int) error
	Update(project *Project) error
}

func NewRepo(conn *gorm.DB) Repo {
	return &repo{conn}
}

func (r *repo) FindWithID(projectID int) (*Project, error) {
	proj := Project{}
	err := r.conn.Where("id = ?", projectID).Find(&proj).Error
	if err != nil {
		return &proj, err
	}

	if proj.ID == 0 {
		return &proj, ErrNoRecord
	}

	return &proj, nil
}

func (r *repo) FindWithUserID(userID int) ([]Project, error) {
	projects := make([]Project, 0)
	err := r.conn.Where("user_id = ?", userID).Find(&projects).Error
	if err != nil {
		return projects, err
	}

	if len(projects) == 0 {
		return projects, ErrNoRecords
	}

	return projects, nil
}

func (r *repo) Save(project *Project) error {
	return r.conn.Create(project).Error
}

func (r *repo) Delete(projectID int) error {
	cmd := r.conn.Delete(Project{ID: projectID})
	if cmd.RowsAffected == 0 {
		return ErrNoRecord
	}

	return cmd.Error
}

func (r *repo) Update(project *Project) error {
	cmd := r.conn.Model(project).Updates(project)
	if cmd.RowsAffected == 0 {
		return ErrNoRecord
	}

	return cmd.Error
}
