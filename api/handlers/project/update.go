package project

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/MihaiBlebea/task-manager/domain"
	"github.com/gorilla/mux"
)

type UpdateRequest struct {
	Title       string `json:"title"`
	Color       string `json:"color"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
}

type UpdateResponse struct {
	ProjectID int    `json:"id,omitempty"`
	Success   bool   `json:"success"`
	Message   string `json:"message,omitempty"`
}

func UpdateHandler(tm domain.TaskManager) http.Handler {
	validate := func(r *http.Request) (int, *UpdateRequest, error) {
		request := UpdateRequest{}

		params := mux.Vars(r)
		id, ok := params["project_id"]
		if ok == false {
			return 0, &request, errors.New("Invalid request param project_id")
		}

		projectID, err := strconv.Atoi(id)
		if err != nil {
			return 0, &request, err
		}

		if projectID == 0 {
			return 0, &request, errors.New("Invalid request param project_id")
		}

		err = json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			return 0, &request, err
		}

		return projectID, &request, nil
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response := UpdateResponse{}

		projectID, req, err := validate(r)
		if err != nil {
			response.Message = err.Error()
			sendResponse(w, response, http.StatusBadRequest)
			return
		}

		err = tm.UpdateProject(1, projectID, req.Title, req.Color, req.Description, req.Icon)
		if err != nil {
			response.Message = err.Error()
			sendResponse(w, response, http.StatusBadRequest)
			return
		}

		response.Success = true
		response.ProjectID = projectID

		sendResponse(w, response, 200)
	})
}
