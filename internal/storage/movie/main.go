package storage

import "github.com/Olegsuus/MoviesGRPC/internal/storage/db"

type MovieStorage struct {
	db *db.MongoStorage
}

func RegisterMovieStorage(mongoStorage *db.MongoStorage) *MovieStorage {
	return &MovieStorage{
		db: mongoStorage,
	}
}
