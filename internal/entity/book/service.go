package book

import (
	"context"
	"fmt"

	"github.com/alexPavlikov/Library-Manegement-System/pkg/logging"
)

type Service struct {
	repository Repository
	logger     *logging.Logger
}

func NewService(logger *logging.Logger, repository Repository) *Service {
	return &Service{
		repository: repository,
		logger:     logger,
	}
}

func (s *Service) Create(ctx context.Context, book *Book) error {
	err := s.repository.Create(ctx, book)
	if err != nil {
		return fmt.Errorf("failed to create book, due to err: %v", err)
	}
	return nil
}
func (s *Service) GetAll(ctx context.Context) ([]Book, error) {
	books, err := s.repository.GetAll(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get all books, due to err: %v", err)
	}
	return books, nil
}
func (s *Service) GetBook(ctx context.Context, uuid string) (Book, error) {
	book, err := s.repository.GetBook(ctx, uuid)
	if err != nil {
		return Book{}, fmt.Errorf("failed to get book, due to err: %v", err)
	}
	return book, nil
}
func (s *Service) Update(ctx context.Context, book *Book) error {
	err := s.repository.Update(ctx, book)
	if err != nil {
		return fmt.Errorf("failed to update book, due to err: %v", err)
	}
	return nil
}
func (s *Service) Delete(ctx context.Context, uuid string) error {
	err := s.repository.Delete(ctx, uuid)
	if err != nil {
		return fmt.Errorf("failed to deleted book, due to err: %v", err)
	}
	return nil
}
