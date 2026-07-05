package auth

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

func (r *PostgresRepository) Create(ctx context.Context, user *User) error {
	user.ID = uuid.NewString()

	query := `
		INSERT INTO users (
			id,
			first_name,
			last_name,
			email,
			password_hash
		)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING created_at, updated_at;
	`

	return r.db.Pool().QueryRow(
		ctx,
		query,
		user.ID,
		user.FirstName,
		user.LastName,
		user.Email,
		user.PasswordHash,
	).Scan(
		&user.CreatedAt,
		&user.UpdatedAt,
	)
}

func (r *PostgresRepository) GetByEmail(ctx context.Context, email string) (*User, error) {
	user := &User{}

	query := `
		SELECT
			id,
			first_name,
			last_name,
			email,
			password_hash,
			created_at,
			updated_at
		FROM users
		WHERE email = $1;
	`

	err := r.db.Pool().QueryRow(ctx, query, email).Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.PasswordHash,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *PostgresRepository) GetByID(ctx context.Context, id string) (*User, error) {
	user := &User{}

	query := `
		SELECT
			id,
			first_name,
			last_name,
			email,
			password_hash,
			created_at,
			updated_at
		FROM users
		WHERE id = $1;
	`

	err := r.db.Pool().QueryRow(ctx, query, id).Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.PasswordHash,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return user, nil
}
