package grpc

import (
	"context"
	"github.com/Olegsuus/MovieProto/gen/models/movie"
	"go.mongodb.org/mongo-driver/bson"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log/slog"
)

func (s *MovieServer) Update(ctx context.Context, req *moviepb.UpdateRequest) (*moviepb.UpdateResponse, error) {
	const op = "grpc.update"

	logger := s.l.With(slog.String("op", op))

	update := bson.M{}
	if req.Movie.Title != "" {
		update["title"] = req.Movie.Title
	}
	if req.Movie.Description != "" {
		update["description"] = req.Movie.Description
	}
	if req.Movie.Year != 0 {
		update["year"] = req.Movie.Year
	}
	if req.Movie.Country != "" {
		update["country"] = req.Movie.Country
	}
	if len(req.Movie.Genres) > 0 {
		update["genres"] = req.Movie.Genres
	}
	if req.Movie.PosterUrl != "" {
		update["poster_url"] = req.Movie.PosterUrl
	}
	if req.Movie.Rating != 0 {
		update["rating"] = req.Movie.Rating
	}

	if len(update) == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "Нет данных для обновления")
	}

	err := s.msP.Update(ctx, req.Movie.Id, update)
	if err != nil {
		logger.Error("Ошибка при обновлении фильма", slog.String("id", req.Movie.Id), slog.Any("error", err))
		return nil, status.Errorf(codes.Internal, "Ошибка при обновлении фильма: %v", err)
	}

	logger.Info("Фильм успешно обновлен", slog.String("id", req.Movie.Id))

	return &moviepb.UpdateResponse{
		Status: true,
	}, nil
}
