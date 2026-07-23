package server

import (
	"net/http"

	"github.com/Manwinder4u/knowledge-ai/backend/internal/ai/chat"
	"github.com/Manwinder4u/knowledge-ai/backend/internal/ai/chunker"
	"github.com/Manwinder4u/knowledge-ai/backend/internal/ai/chunks"
	"github.com/Manwinder4u/knowledge-ai/backend/internal/ai/embedding"
	"github.com/Manwinder4u/knowledge-ai/backend/internal/ai/extractor"
	"github.com/Manwinder4u/knowledge-ai/backend/internal/ai/ingestion"
	"github.com/Manwinder4u/knowledge-ai/backend/internal/ai/llm"
	"github.com/Manwinder4u/knowledge-ai/backend/internal/ai/retrieval"
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

	// Storage
	localStorage := storage.NewLocalStorage(
		cfg.UploadPath,
	)

	// AI
	pdfExtractor := extractor.NewPDFExtractor()
	wordChunker := chunker.NewWordChunker(200, 30)

	// Embedding
	embedder := embedding.NewOllamaEmbedder(
		cfg.OllamaURL,
		cfg.EmbeddingModel,
	)

	chunkRepository := chunks.NewPostgresRepository(db)
	ingestionService := ingestion.NewService(pdfExtractor, wordChunker, chunkRepository, embedder)

	// Retrival
	retrievalService := retrieval.NewService(embedder, chunkRepository)
	llmClient := llm.NewOllamaClient(cfg.OllamaURL, cfg.OllamaChatModel)
	// Chat service
	chatService := chat.NewService(retrievalService, llmClient)
	chatHandler := chat.NewHandler(chatService)

	// Documents
	documentRepo := documents.NewPostgresRepository(db)
	documentService := documents.NewService(
		documentRepo,
		localStorage,
		ingestionService,
	)

	documentHandler := documents.NewHandler(documentService, cfg)

	// Register Auth routes
	auth.RegisterRoutes(s.router, authHandler)
	// Register Documents ROutes
	documents.RegisterRoutes(s.router, authHandler, documentHandler)
	// Register Chat routes
	chat.RegisterRoutes(s.router, authHandler, chatHandler)

	// Existing routes
	s.registerRoutes()
	return s
}

func (s *Server) Start() error {
	return http.ListenAndServe(":"+s.config.Port, s.router)
}
