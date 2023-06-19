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

func (s *Service) CreateBook(ctx context.Context, book *Book) error {
	err := s.repository.CreateBook(ctx, book)
	if err != nil {
		return fmt.Errorf("failed to create book, due to err: %v", err)
	}
	return nil
}
func (s *Service) GetAllBooks(ctx context.Context) ([]Book, error) {
	books, err := s.repository.GetAllBooks(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get all books, due to err: %v", err)
	}
	return books, nil
}

func (s *Service) GetMustPopularBooks(ctx context.Context) ([]Book, error) {
	books, err := s.repository.GetMustPopularBooks(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get most popular books, due to err: %v", err)
	}
	return books, nil
}

func (s *Service) GetBook(ctx context.Context, uuid string) (Book, error) {
	book, err := s.repository.GetOneBook(ctx, uuid)
	if err != nil {
		return Book{}, fmt.Errorf("failed to get book, due to err: %v", err)
	}
	return book, nil
}

func (s *Service) GetAllBooksByGenre(ctx context.Context, uuid string) ([]Book, error) {
	books, err := s.repository.GetAllBooksByGenre(ctx, uuid)
	if err != nil {
		return nil, fmt.Errorf("failed to get all books by genre, due to err: %v", err)
	}
	return books, nil
}

func (s *Service) UpdateBook(ctx context.Context, book *Book) error {
	err := s.repository.UpdateBook(ctx, book)
	if err != nil {
		return fmt.Errorf("failed to update book, due to err: %v", err)
	}
	return nil
}
func (s *Service) DeleteBook(ctx context.Context, uuid string) error {
	err := s.repository.DeleteBook(ctx, uuid)
	if err != nil {
		return fmt.Errorf("failed to deleted book, due to err: %v", err)
	}
	return nil
}

// genre
func (s *Service) GetAllGenres(ctx context.Context) ([]Genre, error) {
	genres, err := s.repository.GetAllGenres(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get all genres, due to err: %s", err)
	}
	return genres, nil
}

func (s *Service) GetGenreByLink(ctx context.Context, link string) (Genre, error) {
	g, err := s.repository.GetGenreByLink(ctx, link)
	if err != nil {
		return Genre{}, fmt.Errorf("failed to get genre by link, due to err: %v", err)
	}
	return g, nil
}

//authors

func (s *Service) FindAllAuthorsByBook(ctx context.Context, id string) ([]Author, error) {
	authors, err := s.repository.FindAllAuthorsByBook(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get all author by book uuid, due to err: %v", err)
	}
	return authors, nil
}
