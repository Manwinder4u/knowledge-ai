package documents

import "context"

type Service struct {
	repo Repository
}

// Constructor
func NewService(repo Repository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) Create(ctx context.Context, userID string, req CreateDocumentRequest) (*Document, error) {

	document := &Document{
		UserID:           userID,
		Title:            req.Title,
		OriginalFilename: req.OriginalFilename,
		Status:           "uploaded",
	}

	if err := s.repo.Create(ctx, document); err != nil {
		return nil, err
	}

	return document, nil
}

func (s *Service) List(
	ctx context.Context,
	userID string,
) ([]Document, error) {

	return s.repo.ListByUser(ctx, userID)
}

func (s *Service) Get(
	ctx context.Context,
	id string,
	userID string,
) (*Document, error) {

	return s.repo.GetByID(ctx, id, userID)
}

func (s *Service) Delete(
	ctx context.Context,
	id string,
	userID string,
) error {

	return s.repo.Delete(ctx, id, userID)
}
