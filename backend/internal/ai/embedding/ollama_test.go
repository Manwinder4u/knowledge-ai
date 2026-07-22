package embedding

import (
	"context"
	"testing"
)

func TestOllamaEmbedder(t *testing.T) {

	embedder := NewOllamaEmbedder(
		"http://localhost:11434",
		"nomic-embed-text",
	)

	vector, err := embedder.Embed(
		context.Background(),
		"Hello from KnowledgeAI",
	)
	if err != nil {
		t.Fatal(err)
	}

	if len(vector) == 0 {
		t.Fatal("embedding is empty")
	}

	t.Log(vector[:10])

	t.Logf("Embedding dimension: %d", len(vector))
}
