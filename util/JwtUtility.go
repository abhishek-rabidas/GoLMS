package util

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/spf13/viper"
	"gomvc/exception"
	"strings"
	"time"
)

var secretKey []byte

func init() {
	viper.SetConfigFile("resources/config.json")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	secretKey = []byte(viper.GetString("secretKey"))
}

func GenerateNewTokenForUser(email string, userType string) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
		"type":  userType,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, _ := token.SignedString(secretKey)

	return tokenString
}

func VerifyToken(tokenString string) (error, *jwt.Token) {

	if strings.TrimSpace(tokenString) == "" {
		return exception.New("Empty token"), nil
	}

	tokenString = tokenString[7:]

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	expirationTime, ok := token.Claims.(jwt.MapClaims)["exp"].(float64)

	if !ok {
		return exception.New("Invalid token format"), nil
	}

	if time.Now().Unix() > int64(expirationTime) {
		return exception.New("Token has expired, relogin"), nil
	}

	if err != nil {
		return err, nil
	}

	if !token.Valid {
		return exception.New("Token is not valid"), nil
	}

	return nil, token
}

func GetUserEmailFromToken(tokenString string) (string, error) {
	err, token := VerifyToken(tokenString)

	if err != nil {
		return "", err
	}

	email := ""

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		email = fmt.Sprint(claims["email"])
	}

	if email == "" {
		return "", exception.New("Invalid payload")
	}

	return email, nil

}
