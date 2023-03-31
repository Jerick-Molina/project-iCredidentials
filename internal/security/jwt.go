package security

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func CreateToken(userId string, website string, secretKey string) (string, error) {

	claims := &jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 10)),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		//iCredidentials website
		Issuer: "iCredidentials.com",
		//Userid
		Audience: []string{userId},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}
	fmt.Println(signedToken)
	if err != nil {
		return "", nil
	}

	return signedToken, nil
}

func ReadToken(tokenString string) (*jwt.RegisteredClaims, error) {
	if tokenString != "" {

		bearer := tokenString[:6]
		tokenString = tokenString[7:]
		if bearer != "Bearer" {
			return nil, fmt.Errorf("not bearer token")
		}
		token, err := jwt.ParseWithClaims(tokenString, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte("pm1tcmj86d6vzst6phnr"), nil

		}, jwt.WithLeeway(5*time.Second))

		if err != nil {
			fmt.Println(err)
			return nil, err

		}

		if claims, ok := token.Claims.(*jwt.RegisteredClaims); ok && token.Valid {
			fmt.Printf("%v ", claims.Issuer)
			return claims, nil
		} else {
			fmt.Println(err)
			return nil, err
		}
	}
	return nil, fmt.Errorf("token empty")
}
