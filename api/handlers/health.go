package handler

import (
	"net/http"

	"github.com/MihaiBlebea/task-manager/api/handlers/utils"
)

func healthEndpoint() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var response struct {
			OK bool `json:"ok"`
		}
		response.OK = true

		utils.SendResponse(w, &response, http.StatusOK)
	})
}
