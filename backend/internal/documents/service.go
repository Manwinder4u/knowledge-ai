package documents

import (
	"context"
	"mime/multipart"

	"github.com/Manwinder4u/knowledge-ai/backend/internal/storage"
)

type Service struct {
	repo    Repository
	storage storage.Storage
}

func NewService(repo Repository, storage storage.Storage) *Service {
	return &Service{
		repo:    repo,
		storage: storage,
	}

}

func (s *Service) Upload(ctx context.Context, userID string, file multipart.File, header *multipart.FileHeader) (*Document, error) {

	path, err := s.storage.Save(ctx, file, header.Filename)
	if err != nil {
		return nil, err
	}

	document := &Document{
		UserID:           userID,
		Title:            header.Filename,
		OriginalFilename: header.Filename,
		StoragePath:      path,
		MimeType:         header.Header.Get("Content-Type"),
		FileSize:         header.Size,
		Status:           "uploaded",
	}

	if err := s.repo.Create(ctx, document); err != nil {
		// If DB operation Fails
		// File remains on Disk forever, Its Automatic Database cleanup
		_ = s.storage.Delete(ctx, document.StoragePath)
		return nil, err
	}

	return document, nil
}

func (s *Service) List(ctx context.Context, userID string) ([]Document, error) {

	return s.repo.ListByUser(ctx, userID)
}

func (s *Service) Get(ctx context.Context, id string, userID string) (*Document, error) {

	return s.repo.GetByID(ctx, id, userID)
}

func (s *Service) Delete(ctx context.Context, id string, userID string) error {

	document, err := s.repo.GetByID(ctx, id, userID)

	if err != nil {
		return err
	}

	// after deleting document ,Also Delting the path from disk as well
	if err := s.storage.Delete(ctx, document.StoragePath); err != nil {
		return err
	}

	return s.repo.Delete(ctx, id, userID)
}
