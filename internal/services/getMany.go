package services

import (
	"context"
	"github.com/Olegsuus/MoviesGRPC/internal/models"
	"log/slog"
)

func (s *MovieService) GetMany(ctx context.Context, sortField string, ascending bool, page, limit int64) ([]*models.Movie, error) {
	const op = "services.getMany"

	logger := s.l.With(slog.String("op", op))

	storageMovies, err := s.msP.GetMany(ctx, sortField, ascending, page, limit)
	if err != nil {
		logger.Error("Ошибка при получении списка фильмов", slog.Any("error", err))
		return nil, err
	}

	var movies []*models.Movie
	for i := range storageMovies {
		movie, err := s.TranslatorToModels(ctx, *storageMovies[i])
		if err != nil {
			logger.Error("Ошибка при переводе фильма", slog.Any("error", err))
			return nil, err
		}
		movies = append(movies, movie)
	}
	logger.Info("Список фильмов успешно получен")
	return movies, nil
}
