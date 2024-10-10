package storage

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log/slog"
)

func (s *MovieStorage) Update(ctx context.Context, id string, update bson.M) error {
	const op = "storage.update"

	logger := slog.With(slog.String("op", op))

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		logger.Error("%s: %s", op, err)
		return fmt.Errorf("некорректный ID фильма: %w", err)
	}

	_, err = s.db.Collection.UpdateOne(ctx, bson.M{"_id": objID}, bson.M{"$set": update})
	if err != nil {
		logger.Error("%s: %s", op, err)
		return fmt.Errorf("ошибка обновления фильма: %w", err)
	}

	logger.Info("Фильм успешно обновлен", slog.String("id", id))

	return nil
}
