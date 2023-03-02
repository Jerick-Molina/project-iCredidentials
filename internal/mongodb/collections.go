package mongodb

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type Collections struct {
	Users    *mongo.Collection
	Settings *mongo.Collection
	Keys     *mongo.Collection
}

func New(db *mongo.Database) *Collections {

	return &Collections{
		Users: db.Collection("Users"),
	}
}
