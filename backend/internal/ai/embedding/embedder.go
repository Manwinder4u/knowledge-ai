package embedding

import "context"

// this package isn’t a business module—it’s an infrastructure adapter around an external API.
// So, we do not need
//.    service.go
//.    repository.go
//.    handler.go

type Embedder interface {
	Embed(ctx context.Context, texts string) ([]float32, error)
}
