package extractor

import "context"

type Extractor interface {
	Extract(ctx context.Context, path string) (string, error)
}
