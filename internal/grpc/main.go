package grpc

import (
	"context"
	moviepb "github.com/Olegsuus/MovieProto/gen/models/movie"
	"github.com/Olegsuus/MoviesGRPC/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"google.golang.org/grpc"
)

type MovieServer struct {
	msP MovieServerProvider
	moviepb.UnimplementedMovieServiceServer
}

type MovieServerProvider interface {
	Add(ctx context.Context, movie *models.Movie) (string, error)
	Remove(ctx context.Context, id string) error
	Update(ctx context.Context, id string, update bson.M) error
	Get(ctx context.Context, id string) (*models.Movie, error)
	GetMany(ctx context.Context, sortField string, ascending bool, page, limit int64) ([]*models.Movie, error)
}

func Register(gRPC *grpc.Server, msP MovieServerProvider) {
	moviepb.RegisterMovieServiceServer(gRPC, &MovieServer{msP: msP})
}
