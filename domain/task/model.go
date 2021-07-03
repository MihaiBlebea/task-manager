package task

import "time"

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
	Priority        string    `json:"priority"`
	Created         time.Time `json:"created"`
}

func New(projectID int, title string) *Task {
	return &Task{ProjectID: projectID, Title: title}
}
