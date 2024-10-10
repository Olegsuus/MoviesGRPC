package services

import (
	"context"
	storage_models "github.com/Olegsuus/MoviesGRPC/internal/storage/models"
	"go.mongodb.org/mongo-driver/bson"
	"log/slog"
)

type MovieService struct {
	msP MovieServiceProvider
	l   *slog.Logger
}

type MovieServiceProvider interface {
	Add(ctx context.Context, storage *storage_models.Movie) (string, error)
	Remove(ctx context.Context, id string) error
	Update(ctx context.Context, id string, update bson.M) error
	Get(ctx context.Context, id string) (*storage_models.Movie, error)
	GetMany(ctx context.Context, sortField string, ascending bool, page, limit int64) ([]*storage_models.Movie, error)
}

func RegisterMovieService(l *slog.Logger, msP MovieServiceProvider) *MovieService {
	return &MovieService{
		l:   l,
		msP: msP,
	}
}
