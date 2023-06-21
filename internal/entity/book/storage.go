package book

import "context"

type Repository interface {
	CreateBook(ctx context.Context, book *Book) error
	GetOneBook(ctx context.Context, uuid string) (Book, error)
	GetAllBooks(ctx context.Context) ([]Book, error)
	GetMustPopularBooks(ctx context.Context) ([]Book, error)
	GetAllBooksByGenre(ctx context.Context, genreUUID string) ([]Book, error)
	GetNewBooks(ctx context.Context) ([]Book, error)
	UpdateBook(ctx context.Context, book *Book) error
	DeleteBook(ctx context.Context, uuid string) error
	FindAllAuthorsByBook(ctx context.Context, uuid string) ([]Author, error)
	FindAllGenreByBook(ctx context.Context, uuid string) ([]Genre, error)
	FindAllAwardsByBook(ctx context.Context, uuid string) ([]Awards, error)
	GetAllGenres(ctx context.Context) ([]Genre, error)
	GetGenreByLink(ctx context.Context, link string) (Genre, error)
}
