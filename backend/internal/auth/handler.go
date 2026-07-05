package auth

import (
	"encoding/json"
	"net/http"

	"github.com/Manwinder4u/knowledge-ai/backend/internal/response"
)

type Handler struct {
	service *Service
}

// Hanlder of Auth -- Which return the Service
func NewHandler(service *Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) Register(w http.ResponseWriter, r *http.Request) {

	var req RegisterRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.Error(w, http.StatusBadRequest, "invalid request body")
		return
	}

	user, err := h.service.Register(r.Context(), req)

	if err != nil {
		switch err {
		case ErrEmailAlreadyExists:
			response.Error(w, http.StatusConflict, err.Error())
		default:
			response.Error(w, http.StatusInternalServerError, "internal server error")
		}
		return
	}

	res := UserResponse{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
	}

	response.JSON(w, http.StatusCreated, res)
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {

	var req LoginRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.Error(w, http.StatusBadRequest, "invalid request body")
		return
	}

	token, err := h.service.Login(r.Context(), req)

	if err != nil {
		switch err {
		case ErrInvalidCredentials:
			response.Error(w, http.StatusUnauthorized, err.Error())
		default:
			response.Error(w, http.StatusInternalServerError, "internal server error")
		}
		return
	}

	res := map[string]string{
		"access_token": token,
	}

	response.JSON(w, http.StatusOK, res)
}

func (h *Handler) Me(w http.ResponseWriter, r *http.Request) {

	userID, ok := r.Context().Value(UserIDContextKey).(string)
	if !ok {
		response.Error(w, http.StatusUnauthorized, "user not found")
		return
	}

	user, err := h.service.Me(r.Context(), userID)
	if err != nil {
		response.Error(w, http.StatusNotFound, "user not found")
		return
	}

	userResponse := UserResponse{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
	}
	response.JSON(w, http.StatusOK, userResponse)
}
