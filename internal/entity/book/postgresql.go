package book

import (
	"context"

	"github.com/alexPavlikov/Library-Manegement-System/pkg/client/postgresql"
	"github.com/alexPavlikov/Library-Manegement-System/pkg/logging"
	"github.com/alexPavlikov/Library-Manegement-System/pkg/utils"
)

type repository struct {
	client postgresql.Client
	logger logging.Logger
}

// Create implements Repository.
func (r *repository) Create(ctx context.Context, book *Book) error {
	query := `
	INSERT INTO 
		public.book 
		(name, genre, year, publishing,	authors, delete)
	VALUES 
		($1, $2, $3, $4, $5)
	RETURNING id
	`
	r.logger.Tracef("SQL Query: %s", utils.FormatQuery(query))

	rows := r.client.QueryRow(ctx, query, &book.Name, &book.Genre, &book.Year, &book.Publishing.UUID, &book.Authors, false)
	err := rows.Scan(book.UUID)
	if err != nil {
		return err
	}
	return nil
}

// GetAll implements Repository.
func (r *repository) GetAll(ctx context.Context) ([]Book, error) {
	query := `
	SELECT id, name, genre, year, publishing, authors
	FROM public.book
	WHERE delete = false;
	`
	var Books []Book
	r.logger.Tracef("SQL Query: %s", utils.FormatQuery(query))

	rows, err := r.client.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var b Book
		err = rows.Scan(&b.UUID, &b.Name, &b.Genre, &b.Year, &b.Publishing.UUID, &b.Authors)
		if err != nil {
			return nil, err
		}
		Books = append(Books, b)
	}
	return Books, nil
}

// GetBook implements Repository.
func (r *repository) GetBook(ctx context.Context, uuid string) (Book, error) {
	query := `
	SELECT id, name, genre, year, publishing, authors
	FROM public.book
	WHERE delete = false AND id = $1
	`

	r.logger.Tracef("SQL Query: %s", utils.FormatQuery(query))

	rows := r.client.QueryRow(ctx, query, uuid)
	var b Book
	err := rows.Scan(&b.UUID, &b.Name, &b.Genre, &b.Year, &b.Publishing.UUID, &b.Authors)
	if err != nil {
		return Book{}, err
	}
	return b, nil
}

// Update implements Repository.
func (r *repository) Update(ctx context.Context, book *Book) error {
	query := `
	UPDATE INTO public.book
	SET name = $1, genre = $2, year = $3, publishing = $4, authors = $5
	WHERE id =  $6
	`

	r.logger.Tracef("SQL Query: %s", utils.FormatQuery(query))

	_, err := r.client.Query(ctx, query, &book.Name, &book.Genre, &book.Year, &book.Publishing.UUID, &book.Authors, &book.UUID)
	if err != nil {
		return err
	}
	return nil
}

// Delete implements Repository.
func (r *repository) Delete(ctx context.Context, uuid string) error {
	query := `
	UPDATE INTO public.book
	SET delete = true
	WHERE id = $1
	`

	r.logger.Tracef("SQL Query: %s", utils.FormatQuery(query))

	_, err := r.client.Query(ctx, query, uuid)
	if err != nil {
		return err
	}
	return nil
}

func NewRepository(client postgresql.Client, logger *logging.Logger) Repository {
	return &repository{
		client: client,
		logger: *logger,
	}
}
