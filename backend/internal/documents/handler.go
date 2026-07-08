package documents

import (
	"net/http"

	"github.com/Manwinder4u/knowledge-ai/backend/internal/auth"
	"github.com/Manwinder4u/knowledge-ai/backend/internal/config"
	"github.com/Manwinder4u/knowledge-ai/backend/internal/response"
	"github.com/go-chi/chi/v5"
)

type Handler struct {
	service *Service
	config  *config.Config
}

func NewHandler(service *Service, cfg *config.Config) *Handler {
	return &Handler{
		service: service,
		config:  cfg,
	}
}

func (h *Handler) Upload(w http.ResponseWriter, r *http.Request) {

	userID, ok := r.Context().Value(auth.UserIDContextKey).(string)
	if !ok {
		response.Error(w, http.StatusUnauthorized, "unauthorized")
		return
	}

	if err := r.ParseMultipartForm(h.config.MaxUploadSize); err != nil {
		response.Error(w, http.StatusBadRequest, "invalid multipart form")
		return
	}

	file, header, err := r.FormFile("file")
	if err != nil {
		response.Error(w, http.StatusBadRequest, "file is required")
		return
	}
	defer file.Close()

	if err := ValidateUpload(header, h.config.MaxUploadSize); err != nil {
		response.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	document, err := h.service.Upload(r.Context(), userID, file, header)

	if err != nil {
		response.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	response.JSON(w, http.StatusCreated, ToResponse(document))
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

	response.JSON(w, http.StatusOK, responses)
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

	response.JSON(w, http.StatusOK, ToResponse(document))
}

func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {

	userID, ok := r.Context().Value(auth.UserIDContextKey).(string)
	if !ok {
		response.Error(w, http.StatusUnauthorized, "unauthorized")
		return
	}

	documentID := chi.URLParam(r, "id")

	if err := h.service.Delete(r.Context(), documentID, userID); err != nil {

		response.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
