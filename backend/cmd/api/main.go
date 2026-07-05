package main

import (
	"log"

	"github.com/Manwinder4u/knowledge-ai/backend/internal/config"
	"github.com/Manwinder4u/knowledge-ai/backend/internal/database"
	"github.com/Manwinder4u/knowledge-ai/backend/internal/logger"
	"github.com/Manwinder4u/knowledge-ai/backend/internal/server"
)

func main() {
	cfg := config.Load()

	db, err := database.New(cfg)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	srv := server.New(cfg, db)

	logger := logger.New()

	logger.Info().
		Str("service", cfg.AppName).
		Str("port", cfg.Port).
		Msg("Server starting")

	logger.Fatal().
		Err(srv.Start()).
		Msg("Failed to connect to database")
}
