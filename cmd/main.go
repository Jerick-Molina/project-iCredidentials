package main

import (
	"context"
	"fmt"
	"projects/iCredidentials/internal/imongo"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// type Settings struct {
// 	//Server
// 	imongo      *imongo.Archive
// 	Error       string
// 	RedirectURI string
// }

func main() {
	params := imongo.Config{
		Username: "",
		Password: "",
		//Host:     "",
		Host:     "",
		Database: "",
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

	t := imongo.NewArchive(client, client.Database(params.Database))

	err = t.StupidTestTx(context.TODO())
	//err = t.TestNoTX(context.TODO())

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully connected and pinged.")
}
