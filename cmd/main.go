package main

import (
	"context"
	"fmt"
	"projects/iCredidentials/cmd/server.go"
	"projects/iCredidentials/internal/database"
)

func main() {

	params := database.Config{
		Username: "JAdmin",
		Password: "Nixon9090%21",
		Host:     "cluster0.d6crvkb.mongodb.net/test",
		Database: "Account",
	}

	client, err := database.RunDatabase(params)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
	fmt.Println("Connected to Databse")

	archive := database.NewArchive(client, params.Database)
	server := server.NewServer(archive)

	server.Start("localhost:8080")
}
