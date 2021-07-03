package user

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
}

func NewRepo(conn *gorm.DB) Repo {
	return &repo{conn}
}

func (r *repo) Save(user *User) error {
	return r.conn.Create(user).Error
}
