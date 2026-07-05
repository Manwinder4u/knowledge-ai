package auth

type Service struct {
	repository Repository
}

// Service of Auth -- Which return the Repository Interface
func NewService(repository Repository) *Service {
	return &Service{
		repository: repository,
	}
}
