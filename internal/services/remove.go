package services

import (
	"context"
	"fmt"
	"log/slog"
)

func (s *MovieService) Remove(ctx context.Context, id string) error {
	const op = "services.remove"

	logger := s.l.With(slog.String("op", op))

	if err := s.msP.Remove(ctx, id); err != nil {
		s.l.Error("Ошибка при удалении фильм", "id", id, "error", err)
		return fmt.Errorf("%s: %s", op, err)
	}

	logger.Info("Фильм успешно удален", "id", id)

	return nil
}
