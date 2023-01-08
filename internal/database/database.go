package database

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type UriConfig interface {
	ConfigUri() string
}
type Config struct {
	Username   string
	Password   string
	Host       string
	Parameters string
	Database   string
}
type Archive struct {
	client *mongo.Client
	db     *mongo.Database
	*collection
}

func (c Config) ConfigURI() string {
	uri := fmt.Sprintf("mongodb+srv://%s:%s@%s", c.Username, c.Password, c.Host)
	if c.Parameters != "" {
		params := fmt.Sprintf("/?%s", c.Parameters)
		return uri + params
	}
	return uri

}
func RunDatabase(conf Config) (*mongo.Client, error) {

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(conf.ConfigURI()))

	if err != nil {
		return nil, err
	}

	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}
	return client, err
}

func NewArchive(client *mongo.Client, db string) *Archive {
	//usr := NewCollection(client.Database(db), "Users")

	return &Archive{client: client, db: client.Database(db)}
}

func (arch *Archive) TestTX() {

}
