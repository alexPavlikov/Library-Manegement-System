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

func (s *Service) GetNewBooks(ctx context.Context) ([]Book, error) {
	books, err := s.repository.GetNewBooks(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get new books, due to err: %v", err)
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

func (s *Service) GetBookByName(ctx context.Context, name string) ([]Book, error) {
	book, err := s.repository.GetBookByName(ctx, name)
	if err != nil {
		return nil, fmt.Errorf("failed to get book by name, due to err: %v", err)
	}
	return book, nil
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

//authors

func (s *Service) FindAllAuthorsByBook(ctx context.Context, id string) ([]Author, error) {
	authors, err := s.repository.FindAllAuthorsByBook(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get all author by book uuid, due to err: %v", err)
	}
	return authors, nil
}

// comments

func (s *Service) CreateCommentForBook(ctx context.Context, comment *Comment) error {
	err := s.repository.CreateCommentForBook(ctx, comment)
	if err != nil {
		return fmt.Errorf("failed to create comment for book, due to err: %v", err)
	}
	return nil
}

func (s *Service) GetAllCommentByBook(ctx context.Context, book_uuid string, user_uuid string) ([]Comment, bool, error) {
	comments, ok, err := s.repository.GetAllComment(ctx, book_uuid, user_uuid)
	if err != nil {
		return nil, false, fmt.Errorf("failed to get all comments by book, due to err; %v", err)
	}
	return comments, ok, nil
}
