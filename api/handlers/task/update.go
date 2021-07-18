package task

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/MihaiBlebea/task-manager/api/handlers/utils"
	"github.com/MihaiBlebea/task-manager/domain"
	"github.com/gorilla/mux"
)

type UpdateRequest struct {
	Title           string `json:"title"`
	Note            string `json:"note"`
	Expire          string `json:"expire"`
	Repeat          bool   `json:"repeat"`
	RepeatDayOfWeek int    `json:"repeat_day_of_week"`
	RepeatTimeOfDay string `json:"repeat_time_of_day"`
	Priority        int    `json:"priority"`
}

type UpdateResponse struct {
	TaskID  int    `json:"id,omitempty"`
	Success bool   `json:"success"`
	Message string `json:"message,omitempty"`
}

func UpdateHandler(tm domain.TaskManager) http.Handler {
	validate := func(r *http.Request) (int, *UpdateRequest, error) {
		request := UpdateRequest{}

		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			return 0, &request, err
		}

		params := mux.Vars(r)
		id, ok := params["task_id"]
		if ok == false {
			return 0, &request, errors.New("Invalid request param task_id")
		}

		taskID, err := strconv.Atoi(id)
		if err != nil {
			return 0, &request, err
		}

		if taskID == 0 {
			return 0, &request, errors.New("Invalid request param task_id")
		}

		if request.Title == "" {
			return 0, &request, errors.New("Invalid request param title")
		}

		return taskID, &request, nil
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response := CreateResponse{}

		userID, err := utils.GetUserIDFromRequest(r)
		if err != nil {
			response.Message = err.Error()
			utils.SendResponse(w, response, http.StatusForbidden)
			return
		}

		taskID, req, err := validate(r)
		if err != nil {
			response.Message = err.Error()
			utils.SendResponse(w, response, http.StatusBadRequest)
			return
		}

		err = tm.UpdateTask(
			userID,
			taskID,
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

		utils.SendResponse(w, response, 200)
	})
}
