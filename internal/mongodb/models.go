package mongodb

type Account struct {
	UserId    string `bson:"_id"`
	Username  string `bson:"User_name"`
	FirstName string `bson:"First_name"`
	LastName  string `bson:"Last_name"`
	Password  string `bson:"Password"`
}

type RegisterdWebsite struct {
	WebsiteId     string `bson:"_id"`
	Owner         string `bson:"User_id"`
	Url           string `bson:"Url"`
	WebsiteSecret string `bson:"Website_Secret_Key"`
}

//STEPS
/*


========
1. User comes from 3rd party website
2. Signs in
3. Gets a access token
4. Everytime a user does a action send in token access if token access is revoked then sign out

*/
type LinkedAccountToWebsite struct {
	UserId    string `bson:"_id"`
	WebsiteId string `bson:"Website_id"`
}
