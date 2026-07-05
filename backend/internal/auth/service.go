package auth

import (
	"context"
	"errors"
)

type Service struct {
	repo Repository
	jwt  *JWTManager
}

// Service of Auth -- Which return the Repository Interface
func NewService(repository Repository, jwt *JWTManager) *Service {
	return &Service{
		repo: repository,
		jwt:  jwt,
	}
}

var (
	ErrEmailAlreadyExists = errors.New("email already exists")
	ErrInvalidCredentials = errors.New("invalid email or password")
)

func (s *Service) Register(ctx context.Context, req RegisterRequest) (*User, error) {

	// Check if email already exists
	existingUser, err := s.repo.GetByEmail(ctx, req.Email)
	if err == nil && existingUser != nil {
		return nil, ErrEmailAlreadyExists
	}

	hash, err := HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	user := &User{
		FirstName:    req.FirstName,
		LastName:     req.LastName,
		Email:        req.Email,
		PasswordHash: hash,
	}

	if err := s.repo.Create(ctx, user); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *Service) Login(ctx context.Context, req LoginRequest) (string, error) {

	user, err := s.repo.GetByEmail(ctx, req.Email)
	if err != nil {
		return "", ErrInvalidCredentials
	}

	if err := CheckPassword(user.PasswordHash, req.Password); err != nil {
		return "", ErrInvalidCredentials
	}

	token, err := s.jwt.GenerateToken(user.ID)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *Service) Me(ctx context.Context, userID string) (*User, error) {
	return s.repo.GetByID(ctx, userID)
}
