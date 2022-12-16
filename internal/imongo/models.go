package imongo

type Account struct {
	UserId    string `bson:"_id"`
	Username  string `bson:"User_name"`
	FirstName string `bson:"First_name"`
	LastName  string `bson:"Last_name"`
	Email     string `bson:"Email"`
	Password  string `bson:"Password"`
	SettingId string `bson:"Setting_id"`
}

//TODO: Abreviate
type Settings struct {
	SettingId       string `bson:"_id"`
	SettingCode     string `bson:"Setting_Code"`
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

// //Maybe if not we do one to many type of deal
// type Keys struct {
// 	KeyId       string
// 	UserId      string
// 	AccessToken string
// }

//Maybe also
// type Error struct {
// 	ErrorMessage string

// }
