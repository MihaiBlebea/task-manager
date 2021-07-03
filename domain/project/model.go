package project

import "time"

type Project struct {
	ID          int       `json:"id"`
	UserID      int       `json:"user_id"`
	Title       string    `json:"title"`
	Color       string    `json:"color"`
	Description string    `json:"description"`
	Icon        string    `json:"icon"`
	Created     time.Time `json:"created" sql:"DEFAULT:CURRENT_TIMESTAMP"`
	Updated     time.Time `json:"updated" sql:"DEFAULT:CURRENT_TIMESTAMP"`
}

func New(userID int, title string) *Project {
	return &Project{UserID: userID, Title: title}
}
