package mongodb

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var testCollections *Collections

func TestMain(m *testing.M) {
	var params Config

	data, err := ioutil.ReadFile("cmd/values.json")
	if err != nil {
		fmt.Println(err)
		return
	}

	err = json.Unmarshal(data, &params)
	if err != nil {
		fmt.Println(err)
	}

	_, err = mongo.Connect(context.TODO(), options.Client().ApplyURI(params.ConfigURI()))

	if err != nil {
		panic(err)
	}

	os.Exit(m.Run())
}
