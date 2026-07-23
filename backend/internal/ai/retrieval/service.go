package retrieval

import (
	"context"
	"fmt"
	"time"

	"github.com/Manwinder4u/knowledge-ai/backend/internal/ai/chunks"
	"github.com/Manwinder4u/knowledge-ai/backend/internal/ai/embedding"
	"github.com/pgvector/pgvector-go"
)

type Service struct {
	embedder embedding.Embedder
	repo     chunks.Repository
}

// Consturctor
func NewService(embedder embedding.Embedder, repo chunks.Repository) *Service {
	return &Service{
		embedder: embedder,
		repo:     repo,
	}
}

func (s *Service) Retrieve(ctx context.Context, question string, limit int) ([]Result, error) {
	start := time.Now()
	vector, err := s.embedder.Embed(ctx, question)
	fmt.Println("Embedding:", time.Since(start))
	if err != nil {
		return nil, err
	}

	start = time.Now()
	foundChunks, err := s.repo.SearchSimilar(ctx, pgvector.NewVector(vector), limit)
	fmt.Println("Search:", time.Since(start))
	if err != nil {
		return nil, err
	}

	results := make([]Result, 0, len(foundChunks))
	for _, chunk := range foundChunks {
		results = append(results, Result{
			Content: chunk.Content,
		})
	}

	return results, nil
}
