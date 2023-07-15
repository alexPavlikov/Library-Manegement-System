package book

import "context"

type Repository interface {
	CreateBook(ctx context.Context, book *Book) error
	GetOneBook(ctx context.Context, uuid string) (Book, error)
	GetAllBooks(ctx context.Context) ([]Book, error)
	GetMustPopularBooks(ctx context.Context) ([]Book, error)
	GetAllBooksByGenre(ctx context.Context, genreUUID string) ([]Book, error)
	GetAllBooksByAuthor(ctx context.Context, authors string) ([]Book, string, error)
	GetNewBooks(ctx context.Context) ([]Book, error)
	GetBookByName(ctx context.Context, name string) ([]Book, error)
	UpdateBook(ctx context.Context, book *Book) error
	DeleteBook(ctx context.Context, uuid string) error
	FindAllAuthorsByBook(ctx context.Context, uuid string) ([]Author, error)
	// FindAllGenreByBook(ctx context.Context, uuid string) ([]Genre, error)
	FindAllAwardsByBook(ctx context.Context, uuid string) ([]Awards, error)
	// GetAllGenres(ctx context.Context) ([]Genre, error)
	// GetGenreByLink(ctx context.Context, link string) (Genre, error)
	CreateCommentForBook(ctx context.Context, comment *Comment) error
	GetAllComment(ctx context.Context, book_uuid string, user_uuid string) ([]Comment, bool, error)

	TestGetAllBooks(ctx context.Context) (map[int][]Book, error)
}
