package services

import (
	"context"
	"fmt"
	"github.com/Olegsuus/MoviesGRPC/internal/models"
	"log/slog"
)

func (s *MovieService) Add(ctx context.Context, movie *models.Movie) (string, error) {
	const op = "services.add"

	logger := s.l.With(slog.String("op", op))

	storageMovie, err := s.TranslatoToStorage(ctx, *movie)
	if err != nil {
		logger.Error("Ошибка при переводе из models в storage уровень")
		return "", fmt.Errorf("%s: %s", op, err)
	}

	id, err := s.msP.Add(ctx, storageMovie)
	if err != nil {
		logger.Error("Ошибка при добавлении фильма", slog.Any("error", err))
		return "", fmt.Errorf("%s: %s", op, err)
	}

	logger.Info("Фильм успешно добавлен", slog.String("id", id))
	return id, nil
}
