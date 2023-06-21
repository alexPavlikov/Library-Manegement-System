package user

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

func (r *repository) CreateUser(ctx context.Context, user *User) error {
	query := `
	INSERT INTO public.user
		(firstname, lastname, age, date_of_birth, gender, login, password_hash)
	VALUES ($1, $2, $3, $4, $5, $6, $7)
	RETURNING id
	`

	r.logger.Tracef("SQL Query: %s", utils.FormatQuery(query))

	fmt.Println(&user.Firstname, &user.Lastname, &user.Age, &user.DateOfBirth, &user.Gender, &user.Login, &user.PasswordHash)

	err := r.client.QueryRow(ctx, query, user.Firstname, user.Lastname, user.Age, user.DateOfBirth, user.Gender, user.Login, user.PasswordHash).Scan(&user.Id)
	if err != nil {
		return err
	}
	return nil
}

func (r *repository) GetUser(ctx context.Context, uuid string) (User, error) {
	query := `
	SELECT id, firstname, lastname, age, date_of_birth, gender, login, password_hash
	FROM public.user
	WHERE deleted = 'false' AND id =  $1;
	`

	r.logger.Tracef("SQL Query: %s", utils.FormatQuery(query))

	rows, err := r.client.Query(ctx, query, uuid)
	if err != nil {
		return User{}, err
	}
	var user User
	err = rows.Scan(&user.Id, &user.Firstname, &user.Lastname, &user.Age, &user.DateOfBirth, &user.Gender, &user.Login, &user.PasswordHash)
	if err != nil {
		return User{}, err
	}
	return user, nil
}

func (r *repository) GetUserByLoginPassword(ctx context.Context, login string, password string) (User, error) {
	query := `
	SELECT id, firstname, lastname, age, date_of_birth, gender, login, password_hash
	FROM public.user
	WHERE deleted = 'false' AND login =  $1 AND password_hash = $2;
	`

	r.logger.Tracef("SQL Query: %s", utils.FormatQuery(query))

	rows, err := r.client.Query(ctx, query, login, password)
	if err != nil {
		return User{}, err
	}
	var user User
	for rows.Next() {
		err = rows.Scan(&user.Id, &user.Firstname, &user.Lastname, &user.Age, &user.DateOfBirth, &user.Gender, &user.Login, &user.PasswordHash)
		if err != nil {
			return User{}, err
		}
	}
	return user, nil
}

func (r *repository) UpdateUser(ctx context.Context, user *User) error {
	query := `
	UPDATE INTO public.user
	SET firstname = $1, lastname = $2, age = $3, date_of_birth = $4, gender = $5, login = $6, password_hash = $7
	WHERE id = $8
	`

	r.logger.Tracef("SQL Query: %s", utils.FormatQuery(query))

	_, err := r.client.Query(ctx, query, &user.Firstname, &user.Lastname, &user.Age, &user.DateOfBirth, &user.Gender, &user.Login, &user.PasswordHash, &user.Id)
	if err != nil {
		return err
	}
	return nil
}

func (r *repository) DeleteUser(ctx context.Context, uuid string) error {
	query := `
	UPDATE INTO public.user
	SET deleted = 'treu'
	WHERE id = $1
	`

	r.logger.Tracef("SQL Query: %s", utils.FormatQuery(query))

	_, err := r.client.Query(ctx, query, uuid)
	if err != nil {
		return err
	}
	return nil

}
