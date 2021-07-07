package user

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/MihaiBlebea/task-manager/api/handlers/utils"
	"github.com/MihaiBlebea/task-manager/domain"
)

type RegisterRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterResponse struct {
	UserID  int    `json:"id,omitempty"`
	Token   string `json:"token,omitempty"`
	Success bool   `json:"success"`
	Message string `json:"message,omitempty"`
}

func RegisterHandler(tm domain.TaskManager) http.Handler {
	validate := func(r *http.Request) (*RegisterRequest, error) {
		request := RegisterRequest{}

		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			return &request, err
		}

		if request.Username == "" {
			return &request, errors.New("Invalid request param username")
		}

		if request.Email == "" {
			return &request, errors.New("Invalid request param email")
		}

		if request.Password == "" {
			return &request, errors.New("Invalid request param password")
		}

		return &request, nil
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response := RegisterResponse{}

		req, err := validate(r)
		if err != nil {
			response.Message = err.Error()
			utils.SendResponse(w, response, http.StatusBadRequest)
			return
		}

		id, token, err := tm.RegisterUser(
			req.Username,
			req.Email,
			req.Password,
		)
		if err != nil {
			response.Message = err.Error()
			utils.SendResponse(w, response, http.StatusBadRequest)
			return
		}

		response.Success = true
		response.UserID = id
		response.Token = token

		utils.SendResponse(w, response, 200)
	})
}
