package task

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/MihaiBlebea/task-manager/api/handlers/utils"
	"github.com/MihaiBlebea/task-manager/domain"
	"github.com/gorilla/mux"
)

type DeleteResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message,omitempty"`
}

func DeleteHandler(tm domain.TaskManager) http.Handler {
	validate := func(r *http.Request) (int, error) {
		params := mux.Vars(r)
		id, ok := params["task_id"]
		if ok == false {
			return 0, errors.New("Invalid request param task_id")
		}

		taskID, err := strconv.Atoi(id)
		if err != nil {
			return 0, err
		}

		if taskID == 0 {
			return 0, errors.New("Invalid request param task_id")
		}

		return taskID, nil
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response := DeleteResponse{}

		taskID, err := validate(r)
		if err != nil {
			response.Message = err.Error()
			utils.SendResponse(w, response, http.StatusBadRequest)
			return
		}

		err = tm.DeleteTask(1, taskID)
		if err != nil {
			response.Message = err.Error()
			utils.SendResponse(w, response, http.StatusBadRequest)
			return
		}

		response.Success = true

		utils.SendResponse(w, response, 200)
	})
}
