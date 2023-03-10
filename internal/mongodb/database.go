package mongodb

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Config struct {
	Username   string
	Password   string
	Host       string
	Parameters string
	Database   string
}

type Database struct {
	client *mongo.Client
	db     *mongo.Database
	*Collections
}

// Creates a uri string format that is required to connect to  MongoDB.
func (config Config) ConfigURI() string {
	uri := fmt.Sprintf("mongodb+srv://%s:%s@%s", config.Username, config.Password, config.Host)
	if config.Parameters != "" {
		params := fmt.Sprintf("/?%s", config.Parameters)
		return uri + params
	}
	return uri
}

func InitDatabase(client *mongo.Client, db *mongo.Database) *Database {

	return &Database{
		client:      client,
		db:          db,
		Collections: New(db),
	}
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

func (db *Database) execTx(ctx context.Context, fn func(sCtx mongo.SessionContext) (interface{}, error)) error {

	session, err := db.client.StartSession()
	if err != nil {
		return err
	}

	_, err = session.WithTransaction(ctx, fn)
	if err != nil {
		session.AbortTransaction(ctx)
		return err
	}

	session.CommitTransaction(ctx)
	return nil
}
