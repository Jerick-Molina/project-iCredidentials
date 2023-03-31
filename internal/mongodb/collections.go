package mongodb

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type Collections struct {
	Users              *mongo.Collection
	Websites           *mongo.Collection
	ValidKeys          *mongo.Collection
	UserLinkedWebsites *mongo.Collection
}

func New(db *mongo.Database) *Collections {

	return &Collections{
		Users:              db.Collection("Users"),
		Websites:           db.Collection("Registerd_Websites"),
		ValidKeys:          db.Collection("Valid_Keys"),
		UserLinkedWebsites: db.Collection("UserLinkedWebsites"),
	}
}
