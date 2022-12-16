package imongo

import (
	"context"
	"errors"
	"projects/iCredidentials/util"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var letters = "ABCDEFGHIJKLMNOQRZTUVWXYZ"
var numbers = "1234567890"
var special = "!?$&%#^"

func (coll *Collections) GetSettingsDefault(ctx context.Context) (Settings, error) {
	filter := bson.D{{"Setting_Code", "Default"}}
	var settings Settings

	result := coll.Settings.FindOne(ctx, filter)

	if result.Err() != nil {
		if result.Err() == mongo.ErrNoDocuments {
			return settings, mongo.ErrNoDocuments
		}
		return settings, result.Err()
	}

	result.Decode(&settings)
	return settings, nil
}

func (coll *Collections) GetSettingsUnique(ctx context.Context, settingCode string) (Settings, error) {
	filter := bson.D{{"Settings_Code", settingCode}}
	var settings Settings

	result := coll.Settings.FindOne(ctx, filter)

	if result.Err() != nil {
		if result.Err() == mongo.ErrNoDocuments {
			return settings, mongo.ErrNoDocuments
		}
		return settings, result.Err()
	}

	result.Decode(&settings)
	return settings, nil
}

type SetSettingsParams struct {
	SettingCode     string `bson:"Settings_Code"`
	OwnerId         string `bson:"Owner_Id"`
	PassSpecialChar bool   `bson:"Req_Pass_Special_Char"`
	PassNumbers     bool   `bson:"Req_Pass_Numbers"`
	PassMinLength   int    `bson:"Pass_Min_Length"`
	UserMinLength   int    `bson:"User_Min_Length"`
	ResetPassword   bool   `bson:"Reset_Password"`
	FindUsername    bool   `bson:"Find_Username"`
	TwoFactor       bool   `bson:"Req_Two_Factor"`
	EmailConfirm    bool   `bson:"Req_Email_Confirm"`
}

func (coll *Collections) SetUniqueSettings(ctx context.Context, args SetSettingsParams) error {

	_, err := coll.Settings.InsertOne(ctx, args)
	if err != nil {
		return err
	}

	return nil
}
func SettingsAccountCreateValidation(user AccountCreateAccountParams, args Settings) error {

	if args.UserMinLength > len(user.Username) {
		// str := fmt.Sprintln("Minimun username length required: %v", args.UserMinLength)
		// return errors.New(str)
	}
	if args.PassMinLength > len(user.Password) {
		// str := fmt.Sprintln("Minimun password length required: %v", args.PassMinLength)
		// return errors.New(str)
	}
	if args.PassSpecialChar == true {
		if err := util.CharFinder(user.Password, special); err != nil {

			return errors.New("Password requires a special chararacter")
		}
	}
	if args.PassNumbers == true {
		if err := util.CharFinder(user.Password, numbers); err != nil {

			return errors.New("Password requires a number")
		}
	}

	return nil
}
