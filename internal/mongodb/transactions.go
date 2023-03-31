package mongodb

import (
	"context"
	"errors"
	"fmt"

	"project/iCredidentials/internal/security"
	"project/iCredidentials/util"

	"go.mongodb.org/mongo-driver/mongo"
)

type SignInParams struct {
	Username  string `json:"Username" bson:"Username"`
	Password  string `json:"Password" bson:"Password"`
	WebsiteId string `json:"Website_Id" bson:"Website_Id"`
}

// ======== Everything that ties with user =============
func (db *Database) SignInTx(ctx context.Context, params SignInParams) (interface{}, error) {

	var token string
	var acc AccountSignInReturn

	var err error

	err = db.execTx(ctx, func(sCtx mongo.SessionContext) (interface{}, error) {

		acc, err = db.SignIn(ctx, params.Username, params.Password)
		if err != nil {
			return nil, err
		}

		web, err := db.UrlWebsiteValidation(ctx, params.WebsiteId)
		if err != nil {
			return nil, err
		}
		token, err = security.CreateToken(acc.UserId, web.Url, web.WebsiteSecret)
		if err != nil {
			return nil, err
		}

		return nil, nil
	})

	if err != nil {
		return nil, err
	}

	x := []any{token, acc}

	return x, nil
}

func (db *Database) CreateAccountTx(ctx context.Context, params CreateAccountParams) (string, error) {
	var token string

	err := db.execTx(ctx, func(sCtx mongo.SessionContext) (interface{}, error) {

		err := db.UsernameDuplicationValidater(ctx, params.Username)
		if err != nil {
			return nil, err
		}

		user_id, err := db.CreateAccount(ctx, params)
		if err != nil {
			return nil, err
		}

		web, err := db.UrlWebsiteValidation(ctx, params.WebsiteId)
		if err != nil {
			return nil, err
		}
		token, err = security.CreateToken(user_id, web.Url, web.WebsiteSecret)
		if err != nil {
			return nil, err
		}
		return nil, nil
	})

	if err != nil {
		return "", err
	}

	return token, err
}

// ======== Everything that ties with registing websites =============
func (db *Database) RegisterWebsiteTx(ctx context.Context, params RegisterWebsiteParams) error {

	var err error

	err = db.execTx(ctx, func(sCtx mongo.SessionContext) (interface{}, error) {

		for {
			params.WebsiteSecret = util.RandomChars(20)
			//Making sure its  nil, meaning it returns no documents found error or a mongodb err (which is not good)
			if err = db.ValidateSecret(ctx, params.WebsiteSecret); err == nil {
				break
			}
			if err != errors.New("invalid secret, website may be unregistered") {
				return nil, err
			}
			fmt.Println(err)
		}
		params.WebsiteSecretEncoded = util.Hasher(params.WebsiteSecret)

		err = db.RegisterWebsite(ctx, params)

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

// ======== Everything that ties with user registered to website =============

// func (db *Database) RegisterLinkForUserToWesbite(ctx context.Context, params UserLinkedToWebsite) error {

// 	err := db.execTx(ctx, func(sCtx mongo.SessionContext) (interface{}, error) {

// 		return nil, nil
// 	})
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }
