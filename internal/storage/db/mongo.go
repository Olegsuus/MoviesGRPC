package db

import (
	"context"
	"fmt"
	"github.com/Olegsuus/MoviesGRPC/internal/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

type MongoStorage struct {
	Client     *mongo.Client
	DataBase   *mongo.Database
	Collection *mongo.Collection
}

func NewMongoStorage(cfg config.Config) (*MongoStorage, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOption := options.Client().ApplyURI(cfg.Mongo.ConnectString)
	client, err := mongo.Connect(ctx, clientOption)
	if err != nil {
		return nil, fmt.Errorf("ошибка подключения к MongoDB: %w", err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("ошибка проверки соединения с MongoDB: %w", err)
	}

	db := client.Database(cfg.Mongo.Database)
	collection := db.Collection(cfg.Mongo.Collections.Movies)

	log.Println("Подключение в бд установлено")

	return &MongoStorage{
		DataBase:   db,
		Client:     client,
		Collection: collection,
	}, nil

}

func (s *MongoStorage) Close() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	return s.Client.Disconnect(ctx)
}
