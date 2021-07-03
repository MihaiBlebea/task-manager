package user

import "time"

type User struct {
	ID       int       `json:"id"`
	Username string    `json:"username"`
	Email    string    `json:"email"`
	Password string    `json:"password"`
	Token    string    `json:"token"`
	Created  time.Time `json:"created"`
}

func New(username, email, password string) *User {
	return &User{Username: username, Email: email, Password: password}
}
