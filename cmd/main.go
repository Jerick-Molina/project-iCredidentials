package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"project/iCredidentials/cmd/server"
	"project/iCredidentials/internal/mongodb"
)

func main() {
	var params mongodb.Config

	data, err := ioutil.ReadFile("cmd/values.json")
	if err != nil {
		fmt.Println(err)
		return
	}

	err = json.Unmarshal(data, &params)
	if err != nil {
		fmt.Println(err)
	}

	client, err := mongodb.RunDatabase(params)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
	database := mongodb.InitDatabase(client, client.Database(params.Database))
	fmt.Println("Successfully connected and pinged.")
	server := server.InitServer(database)

	server.Start("localhost:8080")
}
