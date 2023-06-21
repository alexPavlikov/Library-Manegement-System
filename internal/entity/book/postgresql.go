package book

import (
	"context"
	"time"

	"github.com/alexPavlikov/Library-Manegement-System/pkg/client/postgresql"
	"github.com/alexPavlikov/Library-Manegement-System/pkg/logging"
	"github.com/alexPavlikov/Library-Manegement-System/pkg/utils"
)

type repository struct {
	client postgresql.Client
	logger logging.Logger
}

// Create implements Repository.
func (r *repository) CreateBook(ctx context.Context, book *Book) error {
	query := `
	INSERT INTO 
		public.book 
		(name, photo, year, pages, description, pdf_link, publishing)
	VALUES 
		($1, $2, $3, $4, $5, $6, $7, $8)
	RETURNING id
	`
	r.logger.Tracef("SQL Query: %s", utils.FormatQuery(query))

	rows := r.client.QueryRow(ctx, query, &book.Name, &book.Photo, &book.Year, &book.Pages, &book.Description, &book.PDFLink, &book.Publishing.UUID)
	err := rows.Scan(book.UUID)
	if err != nil {
		return err
	}
	return nil
}

// GetAll implements Repository.
func (r *repository) GetAllBooks(ctx context.Context) ([]Book, error) {
	query := `
	SELECT b.id, b.name, b.photo, b.year, b.pages, b.description, b.pdf_link, b.publishing, p.name
	FROM public.book b
	JOIN public.publishing p ON p.id = b.publishing
	WHERE b.deleted = false;
	`
	var Books []Book
	r.logger.Tracef("SQL Query: %s", utils.FormatQuery(query))

	rows, err := r.client.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var book Book
		err = rows.Scan(&book.UUID, &book.Name, &book.Photo, &book.Year, &book.Pages, &book.Description, &book.PDFLink, &book.Publishing.UUID, &book.Publishing.Name)
		if err != nil {
			return nil, err
		}

		book.Authors, err = r.FindAllAuthorsByBook(ctx, book.UUID)
		if err != nil {
			return Books, err
		}
		book.Genre, err = r.FindAllGenreByBook(ctx, book.UUID)
		if err != nil {
			return Books, err
		}
		// book.Awards, err = r.FindAllAwardsByBook(ctx, book.UUID)
		// if err != nil {
		// 	return Books, err
		// }

		Books = append(Books, book)
	}
	return Books, nil
}

// GetAll implements Repository.
func (r *repository) GetAllBooksByGenre(ctx context.Context, genreUUID string) ([]Book, error) {
	query := `
	SELECT b.id, b.name, b.photo, b.year, b.pages, b.description, b.pdf_link, b.publishing, p.name
	FROM public.book b
	JOIN public.publishing p ON p.id = b.publishing
	WHERE b.deleted = false AND $1 = ANY(SELECT genre_id FROM public.book_genres WHERE book_id = b.id); ;
	`
	var Books []Book
	r.logger.Tracef("SQL Query: %s", utils.FormatQuery(query))

	rows, err := r.client.Query(ctx, query, genreUUID)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var book Book
		err = rows.Scan(&book.UUID, &book.Name, &book.Photo, &book.Year, &book.Pages, &book.Description, &book.PDFLink, &book.Publishing.UUID, &book.Publishing.Name)
		if err != nil {
			return nil, err
		}

		book.Authors, err = r.FindAllAuthorsByBook(ctx, book.UUID)
		if err != nil {
			return Books, err
		}
		book.Genre, err = r.FindAllGenreByBook(ctx, book.UUID)
		if err != nil {
			return Books, err
		}
		// book.Awards, err = r.FindAllAwardsByBook(ctx, book.UUID)
		// if err != nil {
		// 	return Books, err
		// }

		Books = append(Books, book)
	}
	return Books, nil
}

func (r *repository) GetMustPopularBooks(ctx context.Context) ([]Book, error) {
	query := `
	SELECT b.id, b.name, b.photo, b.year, b.pages, b.description, b.pdf_link, b.publishing, p.name, COUNT(br.id) as mycount
	FROM public.book b
	JOIN public.book_rating br ON br.book_id = b.id
	JOIN public.publishing p ON p.id = b.publishing
	GROUP BY  b.id, b.name, b.photo, b.year, b.pages, b.description, b.pdf_link, b.publishing, p.name
	ORDER BY mycount DESC LIMIT 12;
`

	r.logger.Tracef("SQL Query: %s", utils.FormatQuery(query))
	var books []Book
	rows, err := r.client.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var book Book
		var count interface{}
		err = rows.Scan(&book.UUID, &book.Name, &book.Photo, &book.Year, &book.Pages, &book.Description, &book.PDFLink, &book.Publishing.UUID, &book.Publishing.Name, count)
		if err != nil {
			return nil, err
		}
		book.Authors, err = r.FindAllAuthorsByBook(ctx, book.UUID)
		if err != nil {
			return books, err
		}
		book.Genre, err = r.FindAllGenreByBook(ctx, book.UUID)
		if err != nil {
			return books, err
		}
		// book.Awards, err = r.FindAllAwardsByBook(ctx, book.UUID)
		// if err != nil {
		// 	return books, err
		// }
		books = append(books, book)
	}
	return books, nil
}

func (r *repository) GetNewBooks(ctx context.Context) ([]Book, error) {
	query := `
	SELECT b.id, b.name, b.photo, b.year, b.pages, b.description, b.pdf_link, b.publishing, p.name
	FROM public.book b
	JOIN public.publishing p ON p.id = b.publishing
	WHERE b.deleted = 'false' AND b.Year = $1 LIMIT 12;
	`
	year := time.Now().Format("2006")

	rows, err := r.client.Query(ctx, query, year)
	if err != nil {
		return nil, err
	}
	var book Book
	var books []Book
	for rows.Next() {
		err = rows.Scan(&book.UUID, &book.Name, &book.Photo, &book.Year, &book.Pages, &book.Description, &book.PDFLink, &book.Publishing.UUID, &book.Publishing.Name)
		if err != nil {
			return nil, err
		}

		books = append(books, book)
	}
	return books, nil
}

// GetBook implements Repository.
func (r *repository) GetOneBook(ctx context.Context, uuid string) (Book, error) {
	query := `
	SELECT b.id, b.name, b.photo, b.year, b.pages, b.description, b.pdf_link, b.publishing, p.name  FROM public.book b
	JOIN public.publishing p ON p.id = b.publishing
	WHERE b.id = $1 AND b.deleted = 'false';
	`

	r.logger.Tracef("SQL Query: %s", utils.FormatQuery(query))

	rows := r.client.QueryRow(context.TODO(), query, uuid)
	var book Book
	err := rows.Scan(&book.UUID, &book.Name, &book.Photo, &book.Year, &book.Pages, &book.Description, &book.PDFLink, &book.Publishing.UUID, &book.Publishing.Name)
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
	// book.Awards, err = r.FindAllAwardsByBook(ctx, uuid)
	// if err != nil {
	// 	return Book{}, err
	// }

	return book, nil
}

// Update implements Repository.
func (r *repository) UpdateBook(ctx context.Context, book *Book) error {
	query := `
	UPDATE INTO public.book
	SET name = $1, photo = $2, year = $3, pages = $4, description = $6, pdf_link= $7, publishing = $8
	WHERE id =  $9
	`

	r.logger.Tracef("SQL Query: %s", utils.FormatQuery(query))

	_, err := r.client.Query(ctx, query, &book.Name, &book.Photo, &book.Year, &book.Pages, &book.Description, &book.PDFLink, &book.Publishing.UUID, &book.UUID)
	if err != nil {
		return err
	}
	return nil
}

// Delete implements Repository.
func (r *repository) DeleteBook(ctx context.Context, uuid string) error {
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

func (r *repository) FindAllAwardsByBook(ctx context.Context, uuid string) ([]Awards, error) {
	query := `
	SELECT * FROM public.awards
	WHERE book_id = $1
	`

	r.logger.Tracef("SQL Queru: %s", utils.FormatQuery(query))

	rows, err := r.client.Query(ctx, query, uuid)
	if err != nil {
		return nil, err
	}
	var awards []Awards
	for rows.Next() {
		var a Awards
		err = rows.Scan(&a.UUID, &a.Name, &a.YearOfReceipt)
		if err != nil {
			return nil, err
		}
		awards = append(awards, a)
	}

	return awards, nil
}

//Book genre

func (r *repository) GetAllGenres(ctx context.Context) ([]Genre, error) {
	query := `
	SELECT name, link
	FROM public.genre
	WHERE deleted = false
	`

	r.logger.Tracef("SQL Query: %s", utils.FormatQuery(query))

	rows, err := r.client.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	var genre Genre
	var genres []Genre
	for rows.Next() {
		err = rows.Scan(&genre.Name, &genre.Link)
		if err != nil {
			return nil, err
		}
		genres = append(genres, genre)
	}
	return genres, nil
}

func (r *repository) FindAllGenreByBook(ctx context.Context, uuid string) ([]Genre, error) {
	query := `
	SELECT g.name, g.link FROM public.book_genres bg
	JOIN public.genre g ON g.id = bg.genre_id
	WHERE book_id = $1;
	`

	r.logger.Tracef("SQL Query: %s", utils.FormatQuery(query))

	rows, err := r.client.Query(ctx, query, uuid)
	if err != nil {
		return nil, err
	}
	var genres []Genre
	for rows.Next() {
		var g Genre
		err = rows.Scan(&g.Name, &g.Link)
		if err != nil {
			return nil, err
		}
		genres = append(genres, g)
	}
	return genres, nil
}

func (r *repository) GetGenreByLink(ctx context.Context, link string) (Genre, error) {
	query := `
	SELECT 
		id, name, link 
	FROM public.genre
		WHERE link = $1
	`

	r.logger.Tracef("SQL Query: %s", utils.FormatQuery(query))

	rows := r.client.QueryRow(ctx, query, link)
	var genre Genre
	err := rows.Scan(&genre.Id, &genre.Name, &genre.Link)
	if err != nil {
		return Genre{}, err
	}
	return genre, nil
}

func NewRepository(client postgresql.Client, logger *logging.Logger) Repository {
	return &repository{
		client: client,
		logger: *logger,
	}
}
