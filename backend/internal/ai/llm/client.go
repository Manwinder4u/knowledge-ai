package llm

import "context"

type Client interface {
	Generate(ctx context.Context, messages []Message) (string, error)
}
