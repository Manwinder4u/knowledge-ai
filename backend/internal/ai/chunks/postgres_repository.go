package chunks

import (
	"context"

	"github.com/Manwinder4u/knowledge-ai/backend/internal/database"
	"github.com/jackc/pgx/v5"
)

type PostgresRepository struct {
	db *database.Database
}

func NewPostgresRepository(db *database.Database) Repository {
	return &PostgresRepository{
		db: db,
	}
}

func (r *PostgresRepository) CreateMany(
	ctx context.Context,
	chunks []Chunk,
) error {

	if len(chunks) == 0 {
		return nil
	}

	batch := &pgx.Batch{}

	for _, chunk := range chunks {

		batch.Queue(
			`
			INSERT INTO document_chunks
			(
				document_id,
				chunk_index,
				content
			)
			VALUES
			(
				$1,
				$2,
				$3
			)
			`,
			chunk.DocumentID,
			chunk.ChunkIndex,
			chunk.Content,
		)
	}

	results := r.db.Pool().SendBatch(
		ctx,
		batch,
	)

	defer results.Close()

	for range chunks {

		_, err := results.Exec()

		if err != nil {
			return err
		}
	}

	return nil
}

func (r *PostgresRepository) ListByDocument(
	ctx context.Context,
	documentID string,
) ([]Chunk, error) {

	query := `
		SELECT
			id,
			document_id,
			chunk_index,
			content,
			created_at
		FROM document_chunks
		WHERE document_id = $1
		ORDER BY chunk_index
	`

	rows, err := r.db.Pool().Query(
		ctx,
		query,
		documentID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var chunks []Chunk

	for rows.Next() {

		var chunk Chunk

		err := rows.Scan(
			&chunk.ID,
			&chunk.DocumentID,
			&chunk.ChunkIndex,
			&chunk.Content,
			&chunk.CreatedAt,
		)
		if err != nil {
			return nil, err
		}

		chunks = append(chunks, chunk)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return chunks, nil
}
