package imongo

type Account struct {
	UserId    string `bson:"Id"`
	FirstName string `bson:"FirstName"`
	LastName  string `bson:"LastName"`
	Email     string `bson:"Email"`
	Password  string `bson:"Password"`
	SettingId string `bson:"SettingId"`
}

//TODO: Abreviate
type Settings struct {
	SettingId       string
	SettingCode     string
	OwnerId         string
	PassSpecialChar bool
	PassNumbers     bool
	PassLength      bool
	UsrLength       bool
	ResetPassword   bool
	FindUsername    bool
	TwoFactor       bool
}

// //Maybe if not we do one to many type of deal
// type Keys struct {
// 	KeyId       string
// 	UserId      string
// 	AccessToken string
// }
