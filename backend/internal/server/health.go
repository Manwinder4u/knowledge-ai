package server

import (
	"encoding/json"
	"net/http"
)

type HealthResponse struct {
	Status   string `json:"status"`
	Service  string `json:"service"`
	Database string `json:"database"`
}

func (s *Server) health(w http.ResponseWriter, r *http.Request) {
	response := HealthResponse{
		Status:   "ok",
		Service:  s.config.AppName,
		Database: "up",
	}

	if err := s.db.Health(); err != nil {
		response.Status = "error"
		response.Database = "down"
		w.WriteHeader(http.StatusServiceUnavailable)
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
	}
}
