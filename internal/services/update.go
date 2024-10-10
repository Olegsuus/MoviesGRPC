package services

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"log/slog"
)

func (s *MovieService) Update(ctx context.Context, id string, update bson.M) error {
	const op = "op"

	logger := s.l.With(slog.String("op", op))

	err := s.msP.Update(ctx, id, update)
	if err != nil {
		logger.Error("Ошибка при обновлении фильма", "id", id, "error", err)
		return fmt.Errorf("%s: %s", op, err)
	}

	logger.Info("Фильм успешно обновлен", "id", id)

	return nil
}
