package user

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
	Save(task *User) (int, error)
	Update(user *User) error
	FindWithEmail(email string) (*User, error)
	FindWithID(userID int) (*User, error)
}

func NewRepo(conn *gorm.DB) Repo {
	return &repo{conn}
}

func (r *repo) FindWithID(userID int) (*User, error) {
	user := User{}
	err := r.conn.Where("id = ?", userID).Find(&user).Error
	if err != nil {
		return &user, err
	}

	if user.ID == 0 {
		return &user, ErrNoRecord
	}

	return &user, nil
}

func (r *repo) Save(user *User) (int, error) {
	cmd := r.conn.Create(user)
	if cmd.RowsAffected == 0 {
		return 0, ErrNotInserted
	}

	return user.ID, cmd.Error
}

func (r *repo) Update(user *User) error {
	cmd := r.conn.Model(user).Updates(user)
	if cmd.RowsAffected == 0 {
		return ErrNoRecord
	}

	return cmd.Error
}

func (r *repo) FindWithEmail(email string) (*User, error) {
	user := User{}
	err := r.conn.Where("email = ?", email).Find(&user).Error
	if err != nil {
		return &user, err
	}

	if user.ID == 0 {
		return &user, ErrNoRecord
	}

	return &user, nil
}
