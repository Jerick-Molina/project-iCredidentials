package security

import (
	"fmt"
	"project/iCredidentials/util"
	"time"

	"github.com/dgrijalva/jwt-go/v4"
)

var secretKey = []byte("p8cafxzquew4juy1rk9f")

var token = jwt.New(jwt.SigningMethodHS256)

func CreateAccessToken(usrId string, isOwner string, settingType string) (string, error) {

	claims := token.Claims.(jwt.MapClaims)
	claims["iss"] = "localhost" // This URI
	claims["exp"] = time.Now().Add(time.Hour * 10).Unix()
	claims["aud"] = "localhost" //Redirect URI
	claims["client_id"] = usrId
	claims["iat"] = time.Now()
	claims["jti"] = util.RandomChars(10)
	// default user
	// default owner
	//	claims["scope"] =
	tokenString, err := token.SignedString(secretKey)

	if err != nil {

		return "", err
	}

	return tokenString, nil
}
func Create3rdPartyAccessToken(usrId int, role int, companyId int) (string, error) {

	claims := token.Claims.(jwt.MapClaims)
	claims["iss"] = "localhost"
	claims["exp"] = time.Now().Add(time.Hour * 10).Unix()
	claims["aud"] = "localhost"
	//  claims["sub"] = ""
	claims["client_id"] = usrId
	claims["iat"] = time.Now()
	claims["jti"] = util.RandomChars(10)
	//claims["scope"]
	tokenString, err := token.SignedString(secretKey)

	if err != nil {
		//
		return "", nil
	}

	return tokenString, nil
}

func TokenReader(token string) (jwt.MapClaims, error) {
	var err error
	if token != "" {
		bearer := token[:6]
		token = token[7:]
		if bearer != "Bearer" {
			return nil, fmt.Errorf("Not bearer token")
		}
		token, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Token not valid")
			}

			return secretKey, nil
		})

		if err != nil {

			return nil, err
		}
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {

			return claims, nil

		}

		return nil, err
	}

	return nil, err
}
