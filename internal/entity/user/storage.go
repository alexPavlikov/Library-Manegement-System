package user

import "context"

type Repository interface {
	CreateUser(ctx context.Context, user *User) error
	GetUser(ctx context.Context, uuid string) (User, error)
	GetUserByLoginPassword(ctx context.Context, login string, password string) (User, error)
	UpdateUser(ctx context.Context, user *User) error
	DeleteUser(ctx context.Context, uuid string) error
}
