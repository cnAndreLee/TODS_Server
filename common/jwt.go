package common

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

var secret string = "tttttt"

// func ParseToken(token string) (string, error) {
// 	claim, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
// 		return []byte(secret), nil
// 	})
// 	if err != nil {
// 		return "", err
// 	}
// 	return claim.Claims.(jwt.MapClaims)["uid"].(string), err
// }

func ParseToken(token string) (string, error) {
	claim, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return "", err
	}
	return claim.Claims.(jwt.MapClaims)["uid"].(string), err
}

func CreateToken(uid string) (string, error) {
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"uid": uid,
		"exp": time.Now().Add(time.Hour * 150).Unix(),
	})
	token, err := at.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return token, nil
}
