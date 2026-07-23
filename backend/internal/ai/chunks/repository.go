package chunks

import (
	"context"

	"github.com/pgvector/pgvector-go"
)

type Repository interface {
	CreateMany(ctx context.Context, chunks []Chunk) error

	ListByDocument(ctx context.Context, documentID string) ([]Chunk, error)

	SearchSimilar(ctx context.Context, embedding pgvector.Vector, limit int) ([]Chunk, error)
}
