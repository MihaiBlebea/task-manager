package telegram

import (
	"net/http"

	"github.com/MihaiBlebea/task-manager/api/handlers/utils"
	"github.com/MihaiBlebea/task-manager/domain"
	"github.com/MihaiBlebea/task-manager/telegram"
)

type WebhookResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message,omitempty"`
}

func WebhookHandler(tm domain.TaskManager) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response := WebhookResponse{}

		tel := telegram.New(tm)

		var update, err = tel.ParseRequest(r)
		if err != nil {
			response.Message = err.Error()
			utils.SendResponse(w, response, http.StatusBadRequest)
			return
		}

		err = tel.HandleRequest(update)
		if err != nil {
			response.Message = err.Error()
			utils.SendResponse(w, response, http.StatusBadRequest)
			return
		}
	})
}
