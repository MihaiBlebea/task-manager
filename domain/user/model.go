package user

import (
	"time"
)

type User struct {
	ID        int       `json:"id"`
	Username  string    `json:"username"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	ChatID    int       `json:"chat_id" gorm:"unique"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func New(username, firstName, lastName string, chatID int) *User {
	return &User{Username: username, FirstName: firstName, LastName: lastName, ChatID: chatID}
}
