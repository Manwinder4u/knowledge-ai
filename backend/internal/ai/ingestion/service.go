package ingestion

import (
	"context"
	"fmt"

	"github.com/Manwinder4u/knowledge-ai/backend/internal/ai/chunker"
	"github.com/Manwinder4u/knowledge-ai/backend/internal/ai/chunks"
	"github.com/Manwinder4u/knowledge-ai/backend/internal/ai/embedding"
	"github.com/Manwinder4u/knowledge-ai/backend/internal/ai/extractor"
	"github.com/pgvector/pgvector-go"
)

type Service struct {
	extractor  extractor.Extractor
	chunker    chunker.Chunker
	repository chunks.Repository
	embedder   embedding.Embedder
}

func NewService(
	extractor extractor.Extractor,
	chunker chunker.Chunker,
	repository chunks.Repository,
	embedder embedding.Embedder,
) *Service {

	return &Service{
		extractor:  extractor,
		chunker:    chunker,
		repository: repository,
		embedder:   embedder,
	}
}

func (s *Service) Process(ctx context.Context, document *Document) error {

	fmt.Println("1. Process started")
	fmt.Println(document.StoragePath)
	text, err := s.extractor.Extract(
		ctx,
		document.StoragePath,
	)
	fmt.Printf("%q\n", text)
	if err != nil {
		return err
	}

	fmt.Println("2. Extracted text:", len(text))
	chunksText := s.chunker.Chunk(text)

	fmt.Println("3. Chunks generated:", len(chunksText))

	documentChunks := make(
		[]chunks.Chunk,
		0,
		len(chunksText),
	)

	for index, content := range chunksText {

		vector, err := s.embedder.Embed(ctx, content)
		if err != nil {
			fmt.Println("error while embedding", err)
			return fmt.Errorf("generate embedding: %w", err)
		}

		documentChunks = append(
			documentChunks,
			chunks.Chunk{
				DocumentID: document.ID,
				ChunkIndex: index,
				Content:    content,
				Embedding:  pgvector.NewVector(vector),
			},
		)
	}

	fmt.Println("4. Calling CreateMany")

	if err := s.repository.CreateMany(ctx, documentChunks); err != nil {
		return fmt.Errorf("store chunks: %w", err)
	}

	fmt.Println("5. CreateMany returned:", err)

	return err
}
