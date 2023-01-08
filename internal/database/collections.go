package database

import "go.mongodb.org/mongo-driver/mongo"

type Collections interface {
	Save()
}

type collection struct {
	*mongo.Collection
}

const (
	User = "Users"
)

func NewCollection(db *mongo.Database, collec string) Collections {

	switch collec {
	case "Users":
		return &collection{db.Collection(collec)}
	default:
		return nil
	}
))
}

func (collection *collection) Save() {

}
