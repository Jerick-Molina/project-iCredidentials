package imongo

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type AccountCreateAccountParams struct {
	Username  string `bson:"User_name"`
	FirstName string `bson:"FirstName"`
	LastName  string `bson:"LastName"`
	Email     string `bson:"Email"`
	Password  string `bson:"Password"`
	SettingId string `bson:"Setting_Id"`
}

func (coll *Collections) CreateAccount(ctx context.Context, acc AccountCreateAccountParams) (string, error) {
	var userId string
	result, err := coll.Users.InsertOne(ctx, acc)
	if err != nil {
		return "", err
	}

	if oid, ok := result.InsertedID.(primitive.ObjectID); ok {
		userId = oid.Hex()
	}

	return userId, nil
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
	SettingId string `bson:"Setting_Id"`
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

type AccountUserIdReturn struct {
}

func (coll *Collections) FindUser(ctx context.Context, userId string) (AccountUserIdReturn, error) {
	var acc AccountUserIdReturn
	oid, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return acc, nil
	}
	filter := bson.D{{"_id", oid}}

	result := coll.Users.FindOne(ctx, filter)
	if result.Err() != nil {
		return acc, result.Err()
	}

	err = result.Decode(&acc)
	if err != nil {
		return acc, err
	}

	return acc, nil
}
