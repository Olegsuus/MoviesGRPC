package storage

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log/slog"
)

func (s *MovieStorage) Remove(ctx context.Context, id string) error {
	const op = "storage.remove"
	logger := slog.With(slog.String("op", op))

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		logger.Error("%s: %s", op, err)
		return fmt.Errorf("некорректный ID фильма: %w", err)
	}

	result, err := s.db.Collection.DeleteOne(ctx, bson.M{"_id": objID})
	if err != nil {
		logger.Error("%s: %s", op, err)
		return fmt.Errorf("ошибка удаления фильма: %w", err)
	}

	if result.DeletedCount == 0 {
		return fmt.Errorf("фильм с ID %s не найден", id)
	}

	logger.Info("Фильм успешно удален", slog.String("id", id))

	return nil
}
