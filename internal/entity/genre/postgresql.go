package genre

import (
	"context"
	"fmt"

	"github.com/alexPavlikov/Library-Manegement-System/internal/config"
	"github.com/alexPavlikov/Library-Manegement-System/pkg/client/postgresql"
	"github.com/alexPavlikov/Library-Manegement-System/pkg/logging"
)

type repository struct {
	logger *logging.Logger
	client postgresql.Client
}

func NewRepository(logger *logging.Logger, client postgresql.Client) Repository {
	return &repository{
		logger: logger,
		client: client,
	}
}

func GetAllGenres(ctx context.Context) ([]Genre, error) {
	//var logger *logging.Logger
	cfg := config.GetConfig()

	client, err := postgresql.NewClient(context.TODO(), cfg.Storage)
	if err != nil {
		fmt.Println("ERROR!!!", err)
		//logger.Fatalf("failed to get new client postgresql, due to err: %v", err)
	}

	query := `
	SELECT name, link
	FROM public.genre
	WHERE deleted = false
	`

	//logger.Tracef("SQL Query: %s", utils.FormatQuery(query))

	rows, err := client.Query(ctx, query)
	if err != nil {
		fmt.Println("ERROR!!!", err)
		return nil, err
	}
	var genre Genre
	var genres []Genre
	for rows.Next() {
		err = rows.Scan(&genre.Name, &genre.Link)
		if err != nil {
			fmt.Println("ERROR!!!", err)
			return nil, err
		}
		genres = append(genres, genre)
	}
	fmt.Println(genres)
	return genres, nil
}

func FindAllGenreByBook(ctx context.Context, uuid string) ([]Genre, error) {
	//var logger *logging.Logger
	cfg := config.GetConfig()

	client, err := postgresql.NewClient(context.TODO(), cfg.Storage)
	if err != nil {
		//logger.Fatalf("failed to get new client postgresql, due to err: %v", err)
		fmt.Println("ERROR!!!", err)
		return nil, err
	}

	query := `
	SELECT g.name, g.link FROM public.book_genres bg
	JOIN public.genre g ON g.id = bg.genre_id
	WHERE book_id = $1;
	`

	//logger.Tracef("SQL Query: %s", utils.FormatQuery(query))

	rows, err := client.Query(ctx, query, uuid)
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

func GetGenreByLink(ctx context.Context, link string) (Genre, error) {
	//var logger *logging.Logger
	cfg := config.GetConfig()

	client, err := postgresql.NewClient(context.TODO(), cfg.Storage)
	if err != nil {
		//logger.Fatalf("failed to get new client postgresql, due to err: %v", err)
		fmt.Println("ERROR!!!", err)
		return Genre{}, err
	}

	query := `
	SELECT 
		id, name, link 
	FROM public.genre
		WHERE link = $1
	`

	//logger.Tracef("SQL Query: %s", utils.FormatQuery(query))

	rows := client.QueryRow(ctx, query, link)
	var genre Genre
	err = rows.Scan(&genre.Id, &genre.Name, &genre.Link)
	if err != nil {
		return Genre{}, err
	}
	return genre, nil
}
