package mongodb

import (
	"context"
	"errors"
	"fmt"
	"project/iCredidentials/util"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type CreateAccountParams struct {
	FirstName string `json:"First_Name" bson:"First_Name"`
	LastName  string `json:"Last_Name" bson:"Last_Name"`
	Username  string `json:"Username" bson:"Username"`
	Password  string `json:"Password" bson:"Password"`
}

func (coll *Collections) CreateAccount(ctx context.Context, acc CreateAccountParams) (string, error) {
	var userId string
	acc.Password = util.Hasher(acc.Password)
	result, err := coll.Users.InsertOne(ctx, acc)
	if err != nil {
		return "", err
	}

	if oid, ok := result.InsertedID.(primitive.ObjectID); ok {
		userId = oid.Hex()
	}

	return userId, nil
}

type AccountSignInReturn struct {
	UserId    string `bson:"_id"`
	FirstName string `bson:"First_Name"`
	LastName  string `bson:"Last_Name"`
	Username  string `bson:"Username"`
}

// Sign user in if username and password match
func (coll *Collections) SignIn(ctx context.Context, username string, password string) (AccountSignInReturn, error) {
	filter := bson.D{{"Username", username}, {"Password", util.Hasher(password)}}
	fmt.Println(username, password)
	var acc AccountSignInReturn
	results := coll.Users.FindOne(ctx, filter)
	err := results.Decode(&acc)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return AccountSignInReturn{}, fmt.Errorf("InvalidCredidentials")
		}
		return AccountSignInReturn{}, err
	}

	return acc, nil
}

// Checks if username exist
func (coll *Collections) UsernameDuplicationValidater(ctx context.Context, username string) error {
	filter := bson.D{{"Username", username}}

	results := coll.Users.FindOne(ctx, filter)

	if results.Err() != nil {
		if results.Err() == mongo.ErrNoDocuments {
			return nil
		}
		return results.Err()
	}

	return errors.New("username already exist")
}

type AccountUserIdReturn struct {
}

func (coll *Collections) SearchForUser(ctx context.Context, userId string) (AccountUserIdReturn, error) {
	var acc AccountUserIdReturn
	oid, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return acc, err
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

type UserLinkedToWebsite struct {
	UserId    string `json:"User_id" bson:"User_id"`
	WebsiteId string `json:"Website_Id" bson:"Website_Id"`
}

func (coll *Collections) UserLinkedToWebsite(ctx context.Context, userId string, websiteId string) error {
	var link UserLinkedToWebsite
	link.UserId = userId
	link.WebsiteId = websiteId

	_, err := coll.UserLinkedWebsites.InsertOne(ctx, link)
	if err != nil {
		return err
	}

	return nil
}
