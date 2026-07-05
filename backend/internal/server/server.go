package server

import (
	"net/http"

	"github.com/Manwinder4u/knowledge-ai/backend/internal/config"
	"github.com/Manwinder4u/knowledge-ai/backend/internal/database"
	"github.com/go-chi/chi/v5"
)

type Server struct {
	config *config.Config
	db     *database.Database
	router *chi.Mux
}

func New(cfg *config.Config, db *database.Database) *Server {
	s := &Server{
		config: cfg,
		db:     db,
		router: chi.NewRouter(),
	}
	s.registerRoutes()
	return s
}

func (s *Server) Start() error {
	return http.ListenAndServe(":"+s.config.Port, s.router)
}
