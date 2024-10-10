package models

type Movie struct {
	ID          string   `bson:"_id,omitempty"`
	Title       string   `bson:"title"`
	Description string   `bson:"description"`
	Year        int32    `bson:"year"`
	Country     string   `bson:"country"`
	Genres      []string `bson:"genres"`
	PosterURL   string   `bson:"poster_url"`
	Rating      float32  `bson:"rating"`
}
