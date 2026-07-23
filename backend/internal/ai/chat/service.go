package chat

import (
	"context"
	"fmt"
	"time"

	"github.com/Manwinder4u/knowledge-ai/backend/internal/ai/llm"
	"github.com/Manwinder4u/knowledge-ai/backend/internal/ai/prompt"
	"github.com/Manwinder4u/knowledge-ai/backend/internal/ai/retrieval"
)

type Service struct {
	retriever *retrieval.Service
	llm       llm.Client
}

func NewService(retriever *retrieval.Service, llmClient llm.Client) *Service {

	return &Service{
		retriever: retriever,
		llm:       llmClient,
	}
}

func (s *Service) Ask(ctx context.Context, question string) (string, error) {

	results, err := s.retriever.Retrieve(ctx, question, 2)
	if err != nil {
		return "", err
	}

	contexts := make([]string, 0, len(results))

	for _, result := range results {
		contexts = append(contexts, result.Content)
	}

	finalPrompt := prompt.Build(question, contexts)

	fmt.Printf("Prompt length: %d chars\n", len(finalPrompt))
	start := time.Now()
	answer, err := s.llm.Generate(ctx, finalPrompt)
	fmt.Println("LLM:", time.Since(start))
	if err != nil {
		return "", err
	}

	return answer, nil
}
