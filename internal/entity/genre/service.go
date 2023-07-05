package genre

import (
	"github.com/alexPavlikov/Library-Manegement-System/pkg/logging"
)

type Service struct {
	logger     *logging.Logger
	repository Repository
}

func NewService(logger *logging.Logger, repository Repository) *Service {
	return &Service{
		logger:     logger,
		repository: repository,
	}
}

// func (s *Service) GetAllGenres(ctx context.Context) ([]Genre, error) {
// 	genres, err := s.repository.GetAllGenres(ctx)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to get all genres, due to err: %s", err)
// 	}
// 	return genres, nil
// }

// func (s *Service) GetGenreByLink(ctx context.Context, link string) (Genre, error) {
// 	g, err := s.repository.GetGenreByLink(ctx, link)
// 	if err != nil {
// 		return Genre{}, fmt.Errorf("failed to get genre by link, due to err: %v", err)
// 	}
// 	return g, nil
// }

// func (s *Service) FindAllGenreByBook(ctx context.Context, uuid string) ([]Genre, error) {
// 	g, err := s.repository.FindAllGenreByBook(ctx, uuid)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to find genre by book, due to err: %v", err)
// 	}
// 	return g, nil
// }
