package documents

import (
	"context"

	"github.com/Manwinder4u/knowledge-ai/backend/internal/database"
	"github.com/google/uuid"
)

type PostgresRepository struct {
	db *database.Database
}

func NewPostgresRepository(db *database.Database) *PostgresRepository {
	return &PostgresRepository{
		db: db,
	}
}

func (r *PostgresRepository) Create(ctx context.Context, document *Document) error {

	document.ID = uuid.NewString()

	query := `
		INSERT INTO documents ( 
			id,
			user_id,
			title,
			original_filename,
			storage_path,
			mime_type,
			file_size,
			status
		)
		VALUES ($1,$2,$3,$4,$5,$6,$7,$8) RETURNING created_at,updated_at;`

	return r.db.Pool().QueryRow(
		ctx,
		query,
		document.ID,
		document.UserID,
		document.Title,
		document.OriginalFilename,
		document.StoragePath,
		document.MimeType,
		document.FileSize,
		document.Status,
	).Scan(
		&document.CreatedAt,
		&document.UpdatedAt,
	)
}

func (r *PostgresRepository) ListByUser(ctx context.Context, userID string) ([]Document, error) {

	query := `SELECT id,
			user_id,
			title,
			original_filename,
			storage_path,
			mime_type,
			file_size,
			status,
			created_at,
			updated_at
		FROM documents WHERE user_id = $1 ORDER BY created_at DESC;`

	rows, err := r.db.Pool().Query(ctx, query, userID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var documents []Document

	for rows.Next() {

		var doc Document

		err := rows.Scan(
			&doc.ID,
			&doc.UserID,
			&doc.Title,
			&doc.OriginalFilename,
			&doc.StoragePath,
			&doc.MimeType,
			&doc.FileSize,
			&doc.Status,
			&doc.CreatedAt,
			&doc.UpdatedAt,
		)

		if err != nil {
			return nil, err
		}

		documents = append(documents, doc)
	}

	return documents, rows.Err()
}

func (r *PostgresRepository) GetByID(ctx context.Context, id string, userId string) (*Document, error) {

	doc := &Document{}

	query := `SELECT 
			id, 
			user_id,
			title,
			original_filename,
			storage_path,
			mime_type,
			file_size,
			status,
			created_at,
			updated_at
		FROM documents WHERE id = $1 AND user_id = $2; `

	err := r.db.Pool().QueryRow(ctx, query, id, userId).Scan(
		&doc.ID,
		&doc.UserID,
		&doc.Title,
		&doc.OriginalFilename,
		&doc.StoragePath,
		&doc.MimeType,
		&doc.FileSize,
		&doc.Status,
		&doc.CreatedAt,
		&doc.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return doc, nil
}

func (r *PostgresRepository) Delete(ctx context.Context, id string, userID string) error {

	query := `DELETE FROM documents WHERE id = $1 AND user_id = $2;`

	_, err := r.db.Pool().Exec(ctx, query, id, userID)

	return err
}
