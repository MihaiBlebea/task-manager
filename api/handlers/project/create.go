package project

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/MihaiBlebea/task-manager/api/handlers/utils"
	"github.com/MihaiBlebea/task-manager/domain"
)

type CreateRequest struct {
	Title       string `json:"title"`
	Color       string `json:"color"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
}

type CreateResponse struct {
	ProjectID int    `json:"id,omitempty"`
	Success   bool   `json:"success"`
	Message   string `json:"message,omitempty"`
}

func CreateHandler(tm domain.TaskManager) http.Handler {
	validate := func(r *http.Request) (*CreateRequest, error) {
		request := CreateRequest{}

		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			return &request, err
		}

		if request.Title == "" {
			return &request, errors.New("Invalid request param title")
		}

		return &request, nil
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response := CreateResponse{}

		userID, err := utils.GetUserIDFromRequest(r)
		if err != nil {
			response.Message = err.Error()
			utils.SendResponse(w, response, http.StatusForbidden)
			return
		}

		req, err := validate(r)
		if err != nil {
			response.Message = err.Error()
			utils.SendResponse(w, response, http.StatusBadRequest)
			return
		}

		id, err := tm.CreateProject(userID, req.Title, req.Color, req.Description, req.Icon)
		if err != nil {
			response.Message = err.Error()
			utils.SendResponse(w, response, http.StatusBadRequest)
			return
		}

		response.Success = true
		response.ProjectID = id

		utils.SendResponse(w, response, 200)
	})
}
