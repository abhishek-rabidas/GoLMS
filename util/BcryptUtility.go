package util

import (
	"golang.org/x/crypto/bcrypt"
	"gomvc/exception"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)

	if err != nil {
		err = exception.New("Unable to hash the password")
		return "", err
	}

	return string(bytes), nil

}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
