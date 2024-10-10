package storage

import (
	"context"
	"fmt"
	storage_models "github.com/Olegsuus/MoviesGRPC/internal/storage/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log/slog"
)

func (s *MovieStorage) Add(ctx context.Context, movie *storage_models.Movie) (string, error) {
	const op = "storage.add"

	logger := slog.With(slog.String("op", op))

	movie.ID = primitive.NewObjectID().Hex()

	_, err := s.db.Collection.InsertOne(ctx, movie)
	if err != nil {
		logger.Error("Ошибка при добавлении фильма", slog.Any("error", err))
		return "", fmt.Errorf("ошибка при добавлении фильма: %w", err)
	}

	logger.Info("Фильм успешно добавлен", slog.String("id", movie.ID))

	return movie.ID, nil
}
