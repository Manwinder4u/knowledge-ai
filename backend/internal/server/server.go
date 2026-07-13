package server

import (
	"net/http"

	"github.com/Manwinder4u/knowledge-ai/backend/internal/ai/extractor"
	"github.com/Manwinder4u/knowledge-ai/backend/internal/auth"
	"github.com/Manwinder4u/knowledge-ai/backend/internal/config"
	"github.com/Manwinder4u/knowledge-ai/backend/internal/database"
	"github.com/Manwinder4u/knowledge-ai/backend/internal/documents"
	"github.com/Manwinder4u/knowledge-ai/backend/internal/middleware"
	"github.com/Manwinder4u/knowledge-ai/backend/internal/storage"
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

	s.router.Use(middleware.Recovery)
	s.router.Use(middleware.Logger)

	// Authentication dependencies
	repo := auth.NewPostgresRepository(db)

	jwtManager := auth.NewJWTManager(cfg.JWTSecret, cfg.JWTExpiryHours)

	service := auth.NewService(repo, jwtManager)

	authHandler := auth.NewHandler(service)

	// Register Auth routes
	auth.RegisterRoutes(s.router, authHandler)

	// Storage
	localStorage := storage.NewLocalStorage(
		cfg.UploadPath,
	)

	// Wire up documents dependencies
	documentRepo := documents.NewPostgresRepository(db)
	pdfExtractor := extractor.NewPDFExtractor()

	documentService := documents.NewService(documentRepo, localStorage, pdfExtractor)

	documentHandler := documents.NewHandler(documentService, cfg)

	documents.RegisterRoutes(s.router, authHandler, documentHandler)

	// Existing routes
	s.registerRoutes()
	return s
}

func (s *Server) Start() error {
	return http.ListenAndServe(":"+s.config.Port, s.router)
}
