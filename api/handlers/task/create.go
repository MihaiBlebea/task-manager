package task

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/MihaiBlebea/task-manager/api/handlers/utils"
	"github.com/MihaiBlebea/task-manager/domain"
)

type CreateRequest struct {
	SubtaskID       int    `json:"subtask_id"`
	ProjectID       int    `json:"project_id"`
	Title           string `json:"title"`
	Note            string `json:"note"`
	Expire          string `json:"expire"`
	Repeat          bool   `json:"repeat"`
	RepeatDayOfWeek int    `json:"repeat_day_of_week"`
	RepeatTimeOfDay string `json:"repeat_time_of_day"`
	Priority        int    `json:"priority"`
}

type CreateResponse struct {
	TaskID  int    `json:"id,omitempty"`
	Success bool   `json:"success"`
	Message string `json:"message,omitempty"`
}

func CreateHandler(tm domain.TaskManager) http.Handler {
	validate := func(r *http.Request) (*CreateRequest, error) {
		request := CreateRequest{}

		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			return &request, err
		}

		if request.ProjectID == 0 {
			return &request, errors.New("Invalid request param project_id")
		}

		if request.Title == "" {
			return &request, errors.New("Invalid request param title")
		}

		return &request, nil
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response := CreateResponse{}

		req, err := validate(r)
		if err != nil {
			response.Message = err.Error()
			utils.SendResponse(w, response, http.StatusBadRequest)
			return
		}

		id, err := tm.CreateTask(
			1,
			req.SubtaskID,
			req.ProjectID,
			req.Title,
			req.Note,
			req.Expire,
			req.Repeat,
			req.RepeatDayOfWeek,
			req.RepeatTimeOfDay,
			req.Priority,
		)
		if err != nil {
			response.Message = err.Error()
			utils.SendResponse(w, response, http.StatusBadRequest)
			return
		}

		response.Success = true
		response.TaskID = id

		utils.SendResponse(w, response, 200)
	})
}
