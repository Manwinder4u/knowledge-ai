package chat

import (
	"github.com/Manwinder4u/knowledge-ai/backend/internal/auth"
	"github.com/go-chi/chi/v5"
)

func RegisterRoutes(router chi.Router, authHandler *auth.Handler, handler *Handler) {

	router.Group(func(r chi.Router) {
		r.Use(authHandler.AuthMiddleware)
		r.Post("/chat", handler.Ask)
	})
}
