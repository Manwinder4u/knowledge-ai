package chunks

import (
	"time"

	"github.com/pgvector/pgvector-go"
)

type Chunk struct {
	ID         string
	DocumentID string
	ChunkIndex int
	Content    string
	Embedding  pgvector.Vector
	CreatedAt  time.Time
}
