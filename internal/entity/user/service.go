package user

import (
	"context"
	"fmt"

	"github.com/alexPavlikov/Library-Manegement-System/pkg/logging"
)

type Service struct {
	logging    *logging.Logger
	repository Repository
}

func NewService(logging *logging.Logger, repository Repository) *Service {
	return &Service{
		logging:    logging,
		repository: repository,
	}
}

func (s *Service) CreateUser(ctx context.Context, user *User) error {
	err := s.repository.CreateUser(ctx, user)
	if err != nil {
		return fmt.Errorf("failed to create user, due to err: %v", err)
	}
	return nil
}

func (s *Service) GetUser(ctx context.Context, uuid string) (User, error) {
	user, err := s.repository.GetUser(ctx, uuid)
	if err != nil {
		return User{}, fmt.Errorf("failed to get user by uuid, due to err: %v", err)
	}
	return user, nil
}

func (s *Service) GetAuth(ctx context.Context, login string, password string) (User, error) {
	user, err := s.repository.GetUserByLoginPassword(ctx, login, password)
	if err != nil {
		return User{}, fmt.Errorf("failed to find user by login and password, due to err: %v", err)
	}

	if user.Id == "" {
		return User{}, fmt.Errorf("failed to find user by login and password, due to err: %v", err)
	}
	return user, nil
}

func (s *Service) UpdateUser(ctx context.Context, user *User) error {
	err := s.repository.UpdateUser(ctx, user)
	if err != nil {
		return fmt.Errorf("failed to update user, due to err: %v", err)
	}
	return nil
}

func (s *Service) DeleteUser(ctx context.Context, uuid string) error {
	err := s.repository.DeleteUser(ctx, uuid)
	if err != nil {
		return fmt.Errorf("failed to deleted user, due to err: %v", err)
	}
	return nil
}
