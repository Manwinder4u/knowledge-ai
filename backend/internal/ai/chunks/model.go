package chunks

import "time"

type Chunk struct {
	ID         string
	DocumentID string
	ChunkIndex int
	Content    string
	Embedding  []float32
	CreatedAt  time.Time
}
