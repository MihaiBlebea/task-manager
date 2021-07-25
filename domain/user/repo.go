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
	All() ([]User, error)
	DeleteByChatID(chatID int) error
	FindWithChatID(chatID int) (*User, error)
}

func NewRepo(conn *gorm.DB) Repo {
	return &repo{conn}
}

func (r *repo) FindWithChatID(chatID int) (*User, error) {
	user := User{}
	err := r.conn.Where("chat_id = ?", chatID).Find(&user).Error
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

func (r *repo) DeleteByChatID(chatID int) error {
	user := &User{}
	if err := r.conn.Where("chat_id = ?", chatID).Find(user).Error; err != nil {
		return err
	}

	cmd := r.conn.Delete(user)
	if cmd.RowsAffected == 0 {
		return ErrNoRecord
	}

	return cmd.Error
}

func (r *repo) All() ([]User, error) {
	users := make([]User, 0)
	err := r.conn.Find(&users).Error
	if err != nil {
		return []User{}, err
	}

	return users, nil
}
