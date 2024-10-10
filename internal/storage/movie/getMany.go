package storage

import (
	"context"
	"fmt"
	storage_models "github.com/Olegsuus/MoviesGRPC/internal/storage/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log/slog"
)

func (s *MovieStorage) GetMany(ctx context.Context, sortField string, ascending bool, page, limit int64) ([]*storage_models.Movie, error) {
	const op = "storage.getMany"

	logger := slog.With(slog.String("op", op))

	var sortOrder int
	if ascending {
		sortOrder = 1
	} else {
		sortOrder = -1
	}

	findOptions := options.Find()
	findOptions.SetSort(bson.D{{Key: sortField, Value: sortOrder}})
	findOptions.SetSkip((page - 1) * limit)
	findOptions.SetLimit(limit)

	cursor, err := s.db.Collection.Find(ctx, bson.M{}, findOptions)
	if err != nil {
		logger.Error("Ошибка при получении списка фильмов", slog.Any("error", err))
		return nil, fmt.Errorf("ошибка при получении списка фильмов: %w", err)
	}

	defer cursor.Close(ctx)

	var movies []*storage_models.Movie
	for cursor.Next(ctx) {
		var movie storage_models.Movie
		if err = cursor.Decode(&movie); err != nil {
			logger.Error("Ошибка декодирования фильма", slog.Any("error", err))
			return nil, err
		}

		movies = append(movies, &movie)
	}

	if err = cursor.Err(); err != nil {
		logger.Error("Ошибка курсора", slog.Any("error", err))
		return nil, err
	}

	logger.Info("Список фильмов успешно получен")

	return movies, nil

}
