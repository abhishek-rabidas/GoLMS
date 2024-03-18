package util

import (
	"github.com/golang-jwt/jwt"
	"gomvc/exception"
	"time"
)

var secretKey = []byte("a2104991-ecea-4f48-8e3c-612c27afda64")

func GenerateNewTokenForUser(email string, userType string) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
		"type":  userType,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, _ := token.SignedString(secretKey)

	return tokenString
}

func VerifyToken(tokenString string) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	// TODO: Check expiry

	if err != nil {
		return err
	}

	if !token.Valid {
		return exception.New("Token is not valid")
	}

	return nil
}
