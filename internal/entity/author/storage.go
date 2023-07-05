package author

import "context"

type Repository interface {
	GetAuthors(ctx context.Context) ([]Author, error)
	GetAuthor(ctx context.Context, uuid string) (Author, error)
	GetAuthorByName(ctx context.Context, name string) ([]Author, error)
	FindAllBookByAuthors(ctx context.Context, uuid string) ([]Book, error)
	FindBiographyAuthor(ctx context.Context, uuid string) ([]string, error)
}
