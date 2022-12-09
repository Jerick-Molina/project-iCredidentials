package imongo

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
		Users:    db.Collection("Users"),
		Settings: db.Collection("Settings"),
		Keys:     db.Collection("Keys"),
	}
}

// func (coll *Collections) CreateAccount() error {

// 	// _ ,err := coll.Users.InsertOne(context.TODO())
// 	// if err != nil {
// 	// 	return err
// 	// }
// 	return nil
// }
