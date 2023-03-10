package mongodb

import (
	"context"

	"project/iCredidentials/internal/security"

	"go.mongodb.org/mongo-driver/mongo"
)

// ======== Everything that ties with user =============
func (db *Database) SignInTx(ctx context.Context, email string, password string) (interface{}, error) {

	var token string
	var acc AccountSignInReturn
	var err error

	err = db.execTx(ctx, func(sCtx mongo.SessionContext) (interface{}, error) {

		acc, err = db.SignIn(ctx, email, password)
		if err != nil {
			return nil, err
		}

		token, err = security.CreateAccessToken(acc.UserId, "NoUnique", "default")
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

func (db *Database) CreateAccountTx(ctx context.Context, acc AccountCreateAccountParams) (string, error) {
	var token string
	err := db.execTx(ctx, func(sCtx mongo.SessionContext) (interface{}, error) {

		err := db.UsernameDuplicationValidater(ctx, acc.Username)
		if err != nil {
			return nil, err
		}

		user_id, err := db.CreateAccount(ctx, acc)
		if err != nil {
			return nil, err
		}
		token, err = security.CreateAccessToken(user_id, "isOwnerUnique", "default")
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

// ======== Everything that ties with user registered to website =============
