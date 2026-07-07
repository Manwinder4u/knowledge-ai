package documents

import (
	"github.com/Manwinder4u/knowledge-ai/backend/internal/auth"
	"github.com/go-chi/chi/v5"
)

func RegisterRoutes(r chi.Router, authHandler *auth.Handler, handler *Handler) {

	r.Route("/documents", func(r chi.Router) {

		r.Use(authHandler.AuthMiddleware)

		r.Post("/", handler.Create)

		r.Get("/", handler.List)

		r.Get("/{id}", handler.Get)

		r.Delete("/{id}", handler.Delete)
	})
}
