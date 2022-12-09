package imongo

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type AccountCreateAccountParams struct {
	FirstName string `bson:"FirstName"`
	LastName  string `bson:"LastName"`
	Email     string `bson:"Email"`
	Password  string `bson:"Password"`
	SettingId string `bson:"SettingId"`
}

func (coll *Collections) CreateAccount(ctx context.Context, acc AccountCreateAccountParams) error {

	_, err := coll.Users.InsertOne(ctx, acc)
	if err != nil {
		return err
	}
	return nil
}

func (coll *Collections) EmailDuplicateValidation(ctx context.Context, email string) error {
	filter := bson.D{{"Email", email}}

	results := coll.Users.FindOne(ctx, filter)

	if results.Err() != nil {
		if results.Err() == mongo.ErrNoDocuments {
			return nil
		}
		return results.Err()
	}

	return errors.New("Email Already Exist")
}

type AccountSignInReturn struct {
	UserId    string `bson:"_id"`
	FirstName string `bson:"FirstName"`
	LastName  string `bson:"LastName"`
	Email     string `bson:"Email"`
	SettingId string `bson:"SettingId"`
}

func (coll *Collections) SignIn(ctx context.Context, email string, password string) (AccountSignInReturn, error) {
	filter := bson.D{{"Email", email}, {"Password", password}}
	var acc AccountSignInReturn
	results := coll.Users.FindOne(ctx, filter)
	err := results.Decode(&acc)
	if err != nil {
		return AccountSignInReturn{}, err
	}

	return acc, nil
}
