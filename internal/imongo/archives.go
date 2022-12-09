package imongo

import (
	"context"
	"errors"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

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
	*Collections
}

// Creates a uri string format that is required to connect to  MongoDB.
func (config Config) UriConfig() string {
	uri := fmt.Sprintf("mongodb+srv://%s:%s@%s", config.Username, config.Password, config.Host)
	if config.Parameters != "" {
		params := fmt.Sprintf("/?%s", config.Parameters)
		return uri + params
	}
	return uri
}

func NewArchive(client *mongo.Client, db *mongo.Database) *Archive {

	return &Archive{
		client:      client,
		db:          db,
		Collections: New(db),
	}
}

//TODO: Create one database but two collections for transactions
// We would want a way to use whatever collection there is available
// EX:  err :=  archives.db.Users.Insert  || err:= archives.db.Keys.Insert
// we know database is always going to be static. The only thing thats dynamic is the collections
// func execTXwithSession(context Context, collection) {

// }

func (arch *Archive) execTx(ctx context.Context, fn func(sCtx mongo.SessionContext) (interface{}, error)) error {

	session, err := arch.client.StartSession()
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

func (arch *Archive) StupidTestTx(ctx context.Context) error {
	//filters := bson.D{{"Email", bson.D{{"$eq", "emai@gmail.com"}}}}

	err := arch.execTx(ctx, func(sCtx mongo.SessionContext) (interface{}, error) {

		filtersA := bson.D{{"Email", "molina@gmail.com"}}
		filtersB := bson.D{{"Email", "hehe@gmail.com"}}

		var a bson.D
		_, err := arch.Collections.Users.InsertOne(sCtx, filtersA)

		err = errors.New("sheeesh")
		if err != nil {

			return nil, err
		}

		err = arch.Collections.Users.FindOne(sCtx, filtersB).Decode(&a)
		fmt.Println(&a)
		if err != nil {
			return nil, err
		}

		return nil, nil
	})
	if err != nil {
		return err
	}

	return nil
}

func (arch *Archive) TestNoTX(ctx context.Context) error {

	filters := bson.D{{"Email", "emai@gmail.com"}}
	var r bson.D
	err := arch.client.Database("TestDB").Collection("Users").FindOne(ctx, filters).Decode(&r)
	fmt.Println(r)

	if err != nil {
		return err
	}

	return nil
}
