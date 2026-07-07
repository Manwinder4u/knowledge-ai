package documents

import (
	"encoding/json"
	"net/http"

	"github.com/Manwinder4u/knowledge-ai/backend/internal/auth"
	"github.com/Manwinder4u/knowledge-ai/backend/internal/response"
	"github.com/go-chi/chi/v5"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {

	userID, ok := r.Context().Value(auth.UserIDContextKey).(string)
	if !ok {
		response.Error(w, http.StatusUnauthorized, "unauthorized")
		return
	}

	var req CreateDocumentRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.Error(w, http.StatusBadRequest, "invalid request body")
		return
	}

	document, err := h.service.Create(
		r.Context(),
		userID,
		req,
	)

	if err != nil {
		response.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	response.JSON(
		w,
		http.StatusCreated,
		ToResponse(document),
	)
}

func (h *Handler) List(w http.ResponseWriter, r *http.Request) {

	userID, ok := r.Context().Value(auth.UserIDContextKey).(string)
	if !ok {
		response.Error(w, http.StatusUnauthorized, "unauthorized")
		return
	}

	documents, err := h.service.List(
		r.Context(),
		userID,
	)

	if err != nil {
		response.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	responses := make([]DocumentResponse, 0, len(documents))

	for _, document := range documents {
		doc := document
		responses = append(responses, ToResponse(&doc))
	}

	response.JSON(
		w,
		http.StatusOK,
		responses,
	)
}

func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {

	userID, ok := r.Context().Value(auth.UserIDContextKey).(string)
	if !ok {
		response.Error(w, http.StatusUnauthorized, "unauthorized")
		return
	}

	documentID := chi.URLParam(r, "id")

	document, err := h.service.Get(
		r.Context(),
		documentID,
		userID,
	)

	if err != nil {
		response.Error(w, http.StatusNotFound, "document not found")
		return
	}

	response.JSON(
		w,
		http.StatusOK,
		ToResponse(document),
	)
}

func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {

	userID, ok := r.Context().Value(auth.UserIDContextKey).(string)
	if !ok {
		response.Error(w, http.StatusUnauthorized, "unauthorized")
		return
	}

	documentID := chi.URLParam(r, "id")

	if err := h.service.Delete(
		r.Context(),
		documentID,
		userID,
	); err != nil {

		response.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
