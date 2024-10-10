package storage

import (
	"context"
	"fmt"
	storage_models "github.com/Olegsuus/MoviesGRPC/internal/storage/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log/slog"
)

func (s *MovieStorage) Get(ctx context.Context, id string) (*storage_models.Movie, error) {
	const op = "storage.get"

	logger := slog.With(slog.String("op", op))

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		logger.Error("Некорректный ID фильма", slog.String("id", id), slog.Any("error", err))
		return nil, fmt.Errorf("некорректный ID фильма: %w", err)
	}

	var movie storage_models.Movie
	err = s.db.Collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&movie)
	if err != nil {
		logger.Error("Ошибка при получении фильма", slog.String("id", id), slog.Any("error", err))
		return nil, fmt.Errorf("ошибка при получении фильма: %w", err)
	}

	logger.Info("Фильм успешно получен", slog.String("id", id))

	return &movie, nil
}
