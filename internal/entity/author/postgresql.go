package author

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

func NewRepository(client postgresql.Client, logger *logging.Logger) Repository {
	return &repository{
		client: client,
		logger: *logger,
	}
}

func (r *repository) GetAuthors(ctx context.Context) ([]Author, error) {
	query := `
	SELECT id, firstname, lastname, photo, birth_place, age, date_of_birth, date_of_death, gender
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
		err = rows.Scan(&author.UUID, &author.Firstname, &author.Lastname, &author.Photo, &author.BirthPlace, &author.Age, &author.DateOfBirth, &author.DateOfDeath, &author.Gender)
		if err != nil {
			return nil, err
		}
		//author.Books, err = GetAllAuthorsBooks()
		// if err != nil {
		// 	return nil, err
		// }

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
	SELECT id, firstname, lastname, photo, birth_place, age, date_of_birth, date_of_death, gender
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

		err = rows.Scan(&author.UUID, &author.Firstname, &author.Lastname, &author.Photo, &author.BirthPlace, &author.Age, &author.DateOfBirth, &author.DateOfDeath, &author.Gender)
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
	query := `
	SELECT id, firstname, lastname, photo, birth_place, age, date_of_birth, date_of_death, gender
	FROM public.author
	WHERE firstname = $1 OR lastname = $1
	`

	r.logger.Tracef("SQL Query: %s", utils.FormatQuery(query))

	rows, err := r.client.Query(ctx, query, name)
	if err != nil {
		return nil, err
	}
	var author Author
	var authors []Author
	for rows.Next() {
		err = rows.Scan(&author.UUID, &author.Firstname, &author.Lastname, &author.Photo, &author.BirthPlace, &author.Age, &author.DateOfBirth, &author.DateOfDeath, &author.Gender)
		if err != nil {
			return nil, err
		}

		authors = append(authors, author)
	}
	return authors, nil
}
