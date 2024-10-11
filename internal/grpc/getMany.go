package grpc

import (
	"context"
	moviepb "github.com/Olegsuus/MovieProto/gen/models/movie"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log/slog"
)

func (s *MovieServer) GetMany(ctx context.Context, req *moviepb.GetManyRequest) (*moviepb.GetManyResponse, error) {
	const op = "grpc.getMany"

	logger := slog.With(slog.String("op", op))

	sortField := ""
	if req.SortByYear {
		sortField = "year"
	} else if req.SortByTitle {
		sortField = "title"
	} else if req.SortByRating {
		sortField = "rating"
	} else {
		sortField = "title"
	}

	movies, err := s.msP.GetMany(ctx, sortField, req.IsAscending, req.Page, req.Limit)
	if err != nil {
		logger.Error("Ошибка при получении списка фильмов", slog.Any("error", err))
		return nil, status.Errorf(codes.Internal, "Ошибка при получении списка фильмов: %v", err)
	}

	var grpcMovies []*moviepb.Movie
	for _, movie := range movies {
		grpcMovies = append(grpcMovies, &moviepb.Movie{
			Id:          movie.ID,
			Title:       movie.Title,
			Description: movie.Description,
			Year:        movie.Year,
			Country:     movie.Country,
			Genres:      movie.Genres,
			PosterUrl:   movie.PosterURL,
			Rating:      movie.Rating,
		})
	}

	logger.Info("Список фильмов успешно получен")
	return &moviepb.GetManyResponse{
		Movies: grpcMovies,
	}, nil
}
