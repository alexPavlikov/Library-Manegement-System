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
		(name, photo, year, pages, rating, description, pdf_link, publishing)
	VALUES 
		($1, $2, $3, $4, $5, $6, $7, $8)
	RETURNING id
	`
	r.logger.Tracef("SQL Query: %s", utils.FormatQuery(query))

	rows := r.client.QueryRow(ctx, query, &book.Name, &book.Photo, &book.Year, &book.Pages, &book.Rating, &book.Description, &book.PDFLink, &book.Publishing.UUID)
	err := rows.Scan(book.UUID)
	if err != nil {
		return err
	}
	return nil
}

// GetAll implements Repository.
func (r *repository) GetAll(ctx context.Context) ([]Book, error) {
	query := `
	SELECT b.id, b.name, b.photo, b.year, b.pages, b.rating, b.description, b.pdf_link, b.publishing, p.name
	FROM public.book b
	JOIN public.publishing p ON p.id = b.publishing
	WHERE deleted = false;
	`
	var Books []Book
	r.logger.Tracef("SQL Query: %s", utils.FormatQuery(query))

	rows, err := r.client.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var book Book
		err = rows.Scan(&book.UUID, &book.Name, &book.Photo, &book.Year, &book.Pages, &book.Rating, &book.Description, &book.PDFLink, &book.Publishing.UUID, &book.Publishing.Name)
		if err != nil {
			return nil, err
		}
		Books = append(Books, book)
	}
	return Books, nil
}

// GetBook implements Repository.
func (r *repository) GetOne(ctx context.Context, uuid string) (Book, error) {
	query := `
	SELECT b.id, b.name, b.photo, b.year, b.pages, b.rating, b.description, b.pdf_link, b.publishing, p.name  FROM public.book b
	JOIN public.publishing p ON p.id = b.publishing
	WHERE b.id = $1 AND b.deleted = 'false';
	`

	r.logger.Tracef("SQL Query: %s", utils.FormatQuery(query))

	rows := r.client.QueryRow(context.TODO(), query, uuid)
	var book Book
	err := rows.Scan(&book.UUID, &book.Name, &book.Photo, &book.Year, &book.Pages, &book.Rating, &book.Description, &book.PDFLink, &book.Publishing.UUID, &book.Publishing.Name)
	if err != nil {
		return Book{}, err
	}
	book.Authors, err = r.FindAllAuthorsByBook(ctx, uuid)
	if err != nil {
		return Book{}, err
	}
	book.Genre, err = r.FindAllGenreByBook(ctx, uuid)
	if err != nil {
		return Book{}, err
	}
	book.Awards, err = r.FindAllAwardsByBook(ctx, uuid) //dont work
	if err != nil {
		return Book{}, err
	}

	return book, nil
}

// Update implements Repository.
func (r *repository) Update(ctx context.Context, book *Book) error {
	query := `
	UPDATE INTO public.book
	SET name = $1, photo = $2, year = $3, pages = $4, rating = $5, description = $6, pdf_link= $7, publishing = $8
	WHERE id =  $9
	`

	r.logger.Tracef("SQL Query: %s", utils.FormatQuery(query))

	_, err := r.client.Query(ctx, query, &book.Name, &book.Photo, &book.Year, &book.Pages, &book.Rating, &book.Description, &book.PDFLink, &book.Publishing.UUID, &book.UUID)
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

func (r *repository) FindAllAuthorsByBook(ctx context.Context, uuid string) ([]Author, error) {
	query := `
	SELECT 
		a.id, a.firstname, a.lastname, a.photo, a.birth_place, a.age, a.date_of_birth, a.date_of_death, a.gender
	FROM public.book b
	JOIN public.book_authors ba ON ba.book_id = $1
	JOIN public.author a ON a.id = ba.author_id 
	WHERE b.id = $2 AND a.deleted = false;
	`

	r.logger.Tracef("SQL Query: %s", utils.FormatQuery(query))

	rows, err := r.client.Query(ctx, query, uuid, uuid)
	if err != nil {
		return nil, err
	}
	var authors []Author
	for rows.Next() {
		var a Author
		err = rows.Scan(&a.UUID, &a.Firstname, &a.Lastname, &a.Photo, &a.BirthPlace, &a.Age, &a.DateOfBirth, &a.DateOfDeath, &a.Gender)
		if err != nil {
			return nil, err
		}
		authors = append(authors, a)
	}

	return authors, nil
}

func (r *repository) FindAllGenreByBook(ctx context.Context, uuid string) ([]string, error) {
	query := `
	SELECT g.name FROM public.book_genres bg
	JOIN public.genre g ON g.id = bg.genre_id
	WHERE book_id = $1;
	`

	r.logger.Tracef("SQL Query: %s", utils.FormatQuery(query))

	rows, err := r.client.Query(ctx, query, uuid)
	if err != nil {
		return nil, err
	}
	var genres []string
	for rows.Next() {
		var name string
		err = rows.Scan(&name)
		if err != nil {
			return nil, err
		}
		genres = append(genres, name)
	}
	return genres, nil
}

func (r *repository) FindAllAwardsByBook(ctx context.Context, uuid string) ([]Awards, error) {
	return nil, nil
}

func NewRepository(client postgresql.Client, logger *logging.Logger) Repository {
	return &repository{
		client: client,
		logger: *logger,
	}
}
