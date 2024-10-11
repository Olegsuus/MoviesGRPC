package services

import (
	"context"
	"github.com/Olegsuus/MoviesGRPC/internal/models"
	storage_models "github.com/Olegsuus/MoviesGRPC/internal/storage/models"
)

func (s *MovieService) TranslatorToModels(_ context.Context, storageMovie storage_models.Movie) (*models.Movie, error) {

	return &models.Movie{
		ID:          storageMovie.ID.Hex(),
		Title:       storageMovie.Title,
		Description: storageMovie.Description,
		Year:        storageMovie.Year,
		Country:     storageMovie.Country,
		Genres:      storageMovie.Genres,
		PosterURL:   storageMovie.PosterURL,
		Rating:      storageMovie.Rating,
	}, nil
}

func (s *MovieService) TranslatoToStorage(_ context.Context, movie models.Movie) (*storage_models.Movie, error) {
	return &storage_models.Movie{
		Title:       movie.Title,
		Description: movie.Description,
		Year:        movie.Year,
		Country:     movie.Country,
		Genres:      movie.Genres,
		PosterURL:   movie.PosterURL,
		Rating:      movie.Rating,
	}, nil
}
