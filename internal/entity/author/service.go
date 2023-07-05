package author

import (
	"context"
	"fmt"

	"github.com/alexPavlikov/Library-Manegement-System/pkg/logging"
)

type Service struct {
	logger     *logging.Logger
	repository Repository
}

func NewService(logger *logging.Logger, repository Repository) *Service {
	return &Service{
		logger:     logger,
		repository: repository,
	}
}

func (s *Service) GetAllAuthors(ctx context.Context) ([]Author, error) {
	authors, err := s.repository.GetAuthors(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get all authors, due to err: %v", err)
	}
	return authors, nil
}

func (s *Service) GetAuthor(ctx context.Context, uuid string) (Author, error) {
	author, err := s.repository.GetAuthor(ctx, uuid)
	if err != nil {
		return Author{}, fmt.Errorf("failed to get author, due to err: %v", err)
	}
	return author, nil
}

func (s *Service) GetAuthorByName(ctx context.Context, name string) ([]Author, error) {
	authors, err := s.repository.GetAuthorByName(ctx, name)
	if err != nil {
		return nil, fmt.Errorf("failed to get all authors, due to err: %v", err)
	}
	return authors, nil
}

func (s *Service) FindBiography(ctx context.Context, uuid string) ([]string, error) {
	text, err := s.repository.FindBiographyAuthor(ctx, uuid)
	if err != nil {
		return nil, fmt.Errorf("failed to find bio, due to err: %v", err)
	}

	return text, nil
}
