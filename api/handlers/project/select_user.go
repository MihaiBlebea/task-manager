package project

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/MihaiBlebea/task-manager/domain"
	"github.com/MihaiBlebea/task-manager/domain/project"
	"github.com/gorilla/mux"
)

type SelectByUserResponse struct {
	Data    []project.Project `json:"data,omitempty"`
	Success bool              `json:"success"`
	Message string            `json:"message,omitempty"`
}

func SelectUserHandler(tm domain.TaskManager) http.Handler {
	validate := func(r *http.Request) (int, error) {
		params := mux.Vars(r)
		id, ok := params["user_id"]
		if ok == false {
			return 0, errors.New("Invalid request param project_id")
		}

		userID, err := strconv.Atoi(id)
		if err != nil {
			return 0, err
		}

		if userID == 0 {
			return 0, errors.New("Invalid request param project_id")
		}

		return userID, nil
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response := SelectByUserResponse{}

		userID, err := validate(r)
		if err != nil {
			response.Message = err.Error()
			sendResponse(w, response, http.StatusBadRequest)
			return
		}

		projects, err := tm.GetUserProjects(userID)
		if err != nil {
			response.Message = err.Error()
			sendResponse(w, response, http.StatusBadRequest)
			return
		}

		response.Data = projects
		response.Success = true

		sendResponse(w, response, 200)
	})
}