package grpc

import (
	"context"
	moviepb "github.com/Olegsuus/MovieProto/gen/models/movie"
	"github.com/Olegsuus/MoviesGRPC/internal/models"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log/slog"
)

func (s *MovieServer) Add(ctx context.Context, req *moviepb.AddRequest) (*moviepb.AddResponse, error) {
	const op = "grpc.add"

	logger := s.l.With(slog.String("op", op))

	movie := &models.Movie{
		Title:       req.Movie.Title,
		Description: req.Movie.Description,
		Year:        req.Movie.Year,
		Country:     req.Movie.Country,
		Genres:      req.Movie.Genres,
		PosterURL:   req.Movie.PosterUrl,
		Rating:      req.Movie.Rating,
	}

	id, err := s.msP.Add(ctx, movie)
	if err != nil {
		logger.Error("Ошибка при добавлении фильма", slog.Any("error", err))
		return nil, status.Errorf(codes.Internal, "Ошибка при добавлении фильма: %v", err)
	}

	logger.Info("Фильм успешно добавлен", slog.String("id", id))

	return &moviepb.AddResponse{
		Id: id,
	}, nil
}
