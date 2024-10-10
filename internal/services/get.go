package services

import (
	"context"
	"fmt"
	"github.com/Olegsuus/MoviesGRPC/internal/models"
	"log/slog"
)

func (s *MovieService) Get(ctx context.Context, id string) (*models.Movie, error) {
	const op = "services.get"

	logger := s.l.With(slog.String("op", op))

	movieStorage, err := s.msP.Get(ctx, id)
	if err != nil {
		logger.Error("Ошибка при получении фильма", "id", id, "error", err)
		return nil, fmt.Errorf("%s: %s", op, err)
	}

	logger.Info("Фильм успешно получен", "id", id)

	movie, err := s.TranslatorToModels(ctx, *movieStorage)
	if err != nil {
		logger.Error("Ошибка при переводе из storage в models уровень")
		return nil, fmt.Errorf("%s: %s", op, err)
	}

	return movie, nil
}
