package storage

import (
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"

	"github.com/google/uuid"
)

type LocalStorage struct {
	basePath string
}

func NewLocalStorage(basePath string) *LocalStorage {
	return &LocalStorage{
		basePath: basePath,
	}
}

func (s *LocalStorage) Save(ctx context.Context, file multipart.File, filename string) (string, error) {

	// Ensure the uploads directory exists
	if err := os.MkdirAll(s.basePath, 0755); err != nil {
		return "", err
	}

	ext := filepath.Ext(filename)

	newFilename := fmt.Sprintf("%s%s", uuid.NewString(), ext)

	fullPath := filepath.Join(
		s.basePath,
		newFilename,
	)

	dst, err := os.Create(fullPath)
	if err != nil {
		return "", err
	}
	defer dst.Close()

	if _, err := io.Copy(dst, file); err != nil {
		return "", err
	}

	return fullPath, nil
}

func (s *LocalStorage) Delete(ctx context.Context, path string) error {
	_ = ctx
	return os.Remove(path)

}
