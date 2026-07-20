package chunks

import "context"

type Repository interface {
	CreateMany(ctx context.Context, chunks []Chunk) error

	ListByDocument(ctx context.Context, documentID string) ([]Chunk, error)
}
