package database

import (
	"context"
	"fmt"
	"time"

	"github.com/Manwinder4u/knowledge-ai/backend/internal/config"
	"github.com/jackc/pgx/v5/pgxpool"
)

func New(cfg *config.Config) (*Database, error) {
	// Instead of: db.Ping(context.Background())
	// we use: context.WithTimeout(..., 5*time.Second)
	// Imagine PostgreSQL is hanging.
	// Without a timeout, our application might wait indefinitely.
	// With a timeout, startup fails after 5 seconds with a clear error.
	// This is a production best practice.
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	dbPool, err := pgxpool.New(ctx, cfg.DatabaseURL)
	if err != nil {
		return nil, fmt.Errorf("create postgres pool: %w", err)
	}

	if err := dbPool.Ping(ctx); err != nil {
		dbPool.Close()
		return nil, fmt.Errorf("ping postgres: %w", err)
	}

	return &Database{pool: dbPool}, nil
}
