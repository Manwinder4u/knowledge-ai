package chat

import (
	"encoding/json"
	"net/http"

	"github.com/Manwinder4u/knowledge-ai/backend/internal/response"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) Ask(w http.ResponseWriter, r *http.Request) {

	var request AskRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		response.Error(w, http.StatusBadRequest, "invalid request")
		return
	}

	resp, err := h.service.Ask(r.Context(), request.Question)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	answer := AskResponse{Answer: resp}
	response.JSON(w, http.StatusOK, answer)
}
