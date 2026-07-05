package auth

import "github.com/go-chi/chi/v5"

func RegisterRoutes(r chi.Router, handler *Handler) {

	r.Route("/auth", func(r chi.Router) {
		r.Post("/register", handler.Register)
		r.Post("/login", handler.Login)
		r.Group(func(r chi.Router) {
			r.Use(handler.AuthMiddleware)
			r.Get("/me", handler.Me)
		})
	})
}
