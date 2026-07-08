package storage

import (
	"context"
	"mime/multipart"
)

type Storage interface {
	Save(ctx context.Context, file multipart.File, filename string) (string, error)
	Delete(ctx context.Context, path string) error
}
