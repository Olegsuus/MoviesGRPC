package storage_models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Movie struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Title       string             `bson:"title"`
	Description string             `bson:"description"`
	Year        int32              `bson:"year"`
	Country     string             `bson:"country"`
	Genres      []string           `bson:"genres"`
	PosterURL   string             `bson:"poster_url"`
	Rating      float32            `bson:"rating"`
}
