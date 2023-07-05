package author

import (
	"context"
	"fmt"

	"github.com/alexPavlikov/Library-Manegement-System/pkg/client/postgresql"
	"github.com/alexPavlikov/Library-Manegement-System/pkg/logging"
	"github.com/alexPavlikov/Library-Manegement-System/pkg/utils"
)

type repository struct {
	client postgresql.Client
	logger logging.Logger
}

func NewRepository(client postgresql.Client, logger *logging.Logger) Repository {
	return &repository{
		client: client,
		logger: *logger,
	}
}

func (r *repository) GetAuthors(ctx context.Context) ([]Author, error) {
	query := `
	SELECT id, firstname, lastname, patronymic, photo, birth_place, age, date_of_birth, date_of_death, gender
	FROM public.author
	WHERE deleted = 'false';
	`

	r.logger.Tracef("SQL Query: %s", utils.FormatQuery(query))

	rows, err := r.client.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	var authors []Author
	for rows.Next() {
		var author Author
		err = rows.Scan(&author.UUID, &author.Firstname, &author.Patronymic, &author.Lastname, &author.Photo, &author.BirthPlace, &author.Age, &author.DateOfBirth, &author.DateOfDeath, &author.Gender)
		if err != nil {
			return nil, err
		}
		author.Books, err = r.FindAllBookByAuthors(ctx, author.UUID)
		if err != nil {
			return nil, err
		}

		//author.Awards  = GetAllAuthorsAwards()
		// if err != nil {
		// 	return nil, err
		// }

		authors = append(authors, author)
	}
	return authors, nil
}

func (r *repository) GetAuthor(ctx context.Context, uuid string) (Author, error) {
	query := `
	SELECT id, firstname, lastname, patronymic, photo, birth_place, age, date_of_birth, date_of_death, gender
	FROM public.author
	WHERE deleted = 'false' AND id = $1;
	`

	r.logger.Tracef("SQL Query: %s", utils.FormatQuery(query))

	rows, err := r.client.Query(ctx, query, uuid)
	if err != nil {
		return Author{}, err
	}
	var author Author
	for rows.Next() {

		err = rows.Scan(&author.UUID, &author.Firstname, &author.Lastname, &author.Patronymic, &author.Photo, &author.BirthPlace, &author.Age, &author.DateOfBirth, &author.DateOfDeath, &author.Gender)
		if err != nil {
			return Author{}, err
		}
		//author.Books, err = GetAllAuthorsBooks()
		// if err != nil {
		// 	return Author{}, err
		// }

		//author.Awards  = GetAllAuthorsAwards()
		// if err != nil {
		// 	return Author{}, err
		// }

	}
	return author, nil
}

func (r *repository) GetAuthorByName(ctx context.Context, name string) ([]Author, error) {
	newName := "%" + name + "%"
	query := `
	SELECT id, firstname, lastname, patronymic, photo, birth_place, age, date_of_birth, date_of_death, gender
	FROM public.author
	WHERE firstname LIKE $1 OR lastname LIKE $2 AND deleted = 'false'
	`

	r.logger.Tracef("SQL Query: %s", utils.FormatQuery(query))

	var author Author
	var authors []Author

	rows, err := r.client.Query(ctx, query, newName, newName)
	if err != nil {
		fmt.Println("Error ", err)
		return nil, err
	}

	for rows.Next() {
		err = rows.Scan(&author.UUID, &author.Firstname, &author.Lastname, &author.Patronymic, &author.Photo, &author.BirthPlace, &author.Age, &author.DateOfBirth, &author.DateOfDeath, &author.Gender)
		if err != nil {
			fmt.Println("Error ", err)
			return nil, err
		}
		author.Books, err = r.FindAllBookByAuthors(ctx, author.UUID)
		if err != nil {
			fmt.Println("Error ", err)
			return nil, err
		}
		authors = append(authors, author)
	}
	return authors, nil
}

func (r *repository) FindAllBookByAuthors(ctx context.Context, uuid string) ([]Book, error) {
	query := `
	SELECT b.id, b.name, b.photo, b.year, b.pages, b.description, b.pdf_link, b.publishing, p.name  
	FROM public.book b
	JOIN public.publishing p ON p.id = b.publishing
	JOIN public.book_authors ba on ba.author_id = $1 AND ba.book_id = b.id
	WHERE b.deleted = 'false';
	`

	r.logger.Tracef("SQL Query: %s", utils.FormatQuery(query))

	rows, err := r.client.Query(ctx, query, uuid)
	if err != nil {
		return nil, err
	}
	var books []Book
	for rows.Next() {
		var book Book
		err = rows.Scan(&book.UUID, &book.Name, &book.Photo, &book.Year, &book.Pages, &book.Description, &book.PDFLink, &book.Publishing.UUID, &book.Publishing.Name)
		if err != nil {
			return nil, err
		}

		books = append(books, book)

	}
	return books, nil
}

func (r *repository) FindBiographyAuthor(ctx context.Context, uuid string) ([]string, error) {
	query := `
	SELECT born, childhood,	study, beginning_of_creativity,	peak_of_creativity,	death
	FROM public.author_boigraphy
	WHERE author_id = $1
	`

	r.logger.Tracef("SQL Query: %s", utils.FormatQuery(query))

	var a, b, c, d, e, f string
	var text []string

	err := r.client.QueryRow(ctx, query, uuid).Scan(&a, &b, &c, &d, &e, &f)
	if err != nil {
		return nil, err
	}
	text = append(text, a, b, c, d, e, f)

	return text, nil
}
