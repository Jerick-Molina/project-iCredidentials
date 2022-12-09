package imongo

import (
	"context"
	"os"
	"testing"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var testCollections *Collections

func TestMain(m *testing.M) {
	params := Config{
		Username: "JAdmin",
		Password: "Nixon9090%21",
		//Host:     "192.168.3.139",
		Host:     "cluster0.d6crvkb.mongodb.net/test",
		Database: "iCredidentials",
	}
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(params.UriConfig()))

	if err != nil {
		panic(err)
	}
	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}

	testCollections = New(client.Database(params.Database))

	os.Exit(m.Run())
}
