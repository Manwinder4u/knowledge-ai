package server

import (
	"github.com/Manwinder4u/knowledge-ai/backend/internals/config"
	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog"
)

type Server struct {
	Router *chi.Mux

	Config *config.Config

	Logger zerolog.Logger
}
