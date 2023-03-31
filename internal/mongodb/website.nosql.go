package mongodb

import (
	"context"
	"errors"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type WebsiteParams struct {
	Owner         string `json:"User_id" bson:"User_id"`
	Url           string `bson:"Url"`
	WebsiteSecret string `bson:"Website_Secret_Key"`
}

// User is able to register website
func (coll *Collections) RegisterWebsite(ctx context.Context, params WebsiteParams) error {

	_, err := coll.Websites.InsertOne(ctx, params)
	if err != nil {
		return err
	}

	return nil
}

func (coll *Collections) GetRegisteredWebsites(ctx context.Context, usrId string) ([]RegisterdWebsite, error) {
	var rWebs []RegisterdWebsite
	filter := bson.D{{"User_id", usrId}}

	result, err := coll.Websites.Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	for result.Next(ctx) {
		var web RegisterdWebsite
		result.Decode(&web)

		rWebs = append(rWebs, web)
	}

	return rWebs, nil
}

// Validates secret from token, check if it matches a registered website if so it allows access
// I can also use it to check if secret already exist so it wont create duplicates
func (coll *Collections) ValidateSecret(ctx context.Context, eSecret string) error {
	filter := bson.D{{"Website_Secret_Encoded", eSecret}}

	result := coll.Websites.FindOne(ctx, filter)

	if result.Err() != nil {

		if result.Err() == mongo.ErrNoDocuments {
			return nil
		}
		return result.Err()
	}
	//No error means secret was valid
	return errors.New("invalid secret, website may be unregistered")
}

// Deletes registered website. If that website tries to make a call, ValidateSecret function will deny access.
func (coll *Collections) DeleteRegisteredWebsite(ctx context.Context, secret string) error {
	filter := bson.D{{"Website_Website_Secret_Encoded", secret}}

	result := coll.Websites.FindOneAndDelete(ctx, filter)
	if result.Err() != nil {
		return result.Err()
	}
	return nil
}

// If it has a redirect use that secret key if not use default.
func (coll *Collections) UrlWebsiteValidation(ctx context.Context, websiteId string) (WebsiteParams, error) {

	var rWeb WebsiteParams

	if websiteId == "" {

		oid, err := primitive.ObjectIDFromHex("64138cc0feb7c595995a74e2")
		if err != nil {
			fmt.Println(err)
			return rWeb, err
		}
		filter := bson.D{{"_id", oid}}

		result := coll.Websites.FindOne(ctx, filter)

		if result.Err() != nil {

			return rWeb, result.Err()
		}

		result.Decode(&rWeb)

	} else if websiteId != "" {

		oid, err := primitive.ObjectIDFromHex(websiteId)
		if err != nil {
			fmt.Println(err)
			return rWeb, err
		}
		filter := bson.D{{"_id", oid}}

		result := coll.Websites.FindOne(ctx, filter)

		if result.Err() != nil {
			if result.Err() == mongo.ErrNoDocuments {
				return rWeb, fmt.Errorf("invalid secret or url, results do not match. website may not be registered")
			}
			return rWeb, result.Err()
		}

		result.Decode(&rWeb)
	} else {
		//ERROR: one or the other is empty so theres no reason to check if valid. must have both url and secret inputs
		return rWeb, fmt.Errorf("error: url or secret is empty cannot validate")
	}

	return rWeb, nil
}
