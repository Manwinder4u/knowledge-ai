package documents

import "context"

type Repository interface {
	Create(ctx context.Context, document *Document) error

	ListByUser(ctx context.Context, userID string) ([]Document, error)

	GetByID(ctx context.Context, id string, userId string) (*Document, error)

	Delete(ctx context.Context, id string, userId string) error
}
