package mongodb

type Account struct {
	UserId    string `bson:"_id"`
	Username  string `bson:"User_name"`
	FirstName string `bson:"First_name"`
	LastName  string `bson:"Last_name"`
	Password  string `bson:"Password"`
}

type RegisterdWebsite struct {
	WebsiteId         string `bson:"_id"`
	Owner             string `bson:"User_id"`
	WebsiteRedirect   string `bson:"Website_Redirect"`
	WebsiteSpecialKey string `bson:"Website_Special_Key"`
}

type LinkedAccountToWebsite struct {
	UserId    string `bson:"_id"`
	WebsiteId string `bson:"Website_id"`
}
