package project

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/MihaiBlebea/task-manager/api/handlers/utils"
	"github.com/MihaiBlebea/task-manager/domain"
	"github.com/MihaiBlebea/task-manager/domain/project"
	"github.com/gorilla/mux"
)

type SelectResponse struct {
	Data    *project.Project `json:"data,omitempty"`
	Success bool             `json:"success"`
	Message string           `json:"message,omitempty"`
}

func SelectIdHandler(tm domain.TaskManager) http.Handler {
	validate := func(r *http.Request) (int, error) {
		params := mux.Vars(r)
		id, ok := params["project_id"]
		if ok == false {
			return 0, errors.New("Invalid request param project_id")
		}

		projectID, err := strconv.Atoi(id)
		if err != nil {
			return 0, err
		}

		if projectID == 0 {
			return 0, errors.New("Invalid request param project_id")
		}

		return projectID, nil
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response := SelectResponse{}

		userID, err := utils.GetUserIDFromRequest(r)
		if err != nil {
			response.Message = err.Error()
			utils.SendResponse(w, response, http.StatusForbidden)
			return
		}

		projectID, err := validate(r)
		if err != nil {
			response.Message = err.Error()
			utils.SendResponse(w, response, http.StatusBadRequest)
			return
		}

		proj, err := tm.GetProject(userID, projectID)
		if err != nil {
			response.Message = err.Error()
			utils.SendResponse(w, response, http.StatusBadRequest)
			return
		}

		response.Data = proj
		response.Success = true

		utils.SendResponse(w, response, 200)
	})
}
