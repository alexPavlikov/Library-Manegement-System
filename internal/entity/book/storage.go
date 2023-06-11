package book

import "context"

type Repository interface {
	Create(ctx context.Context, book *Book) error
	GetBook(ctx context.Context, uuid string) (Book, error)
	GetAll(ctx context.Context) ([]Book, error)
	Update(ctx context.Context, book *Book) error
	Delete(ctx context.Context, uuid string) error
	FindAllAuthorsByBook(ctx context.Context, uuid string) ([]Author, error)
	FindAllGenreByBook(ctx context.Context, uuid string) ([]string, error)
}
