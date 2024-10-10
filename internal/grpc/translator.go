package grpc

import (
	"context"
	"github.com/Olegsuus/MovieProto/gen/models/movie"
	"github.com/Olegsuus/MoviesGRPC/internal/models"
)

func (s *MovieServer) TranslatorToGrpc(_ context.Context, movie models.Movie) (*moviepb.GetResponse, error) {
	return &moviepb.GetResponse{
		Movie: &moviepb.Movie{
			Id:          movie.ID,
			Title:       movie.Title,
			Description: movie.Description,
			Year:        movie.Year,
			Country:     movie.Country,
			Genres:      movie.Genres,
			PosterUrl:   movie.PosterURL,
			Rating:      movie.Rating,
		},
	}, nil
}
