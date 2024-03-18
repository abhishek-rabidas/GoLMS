package model

import (
	"encoding/json"
	"github.com/google/uuid"
	"io"
	"log"
)

type User struct {
	UID      uuid.UUID `json:"uid"`
	Name     string    `json:"name"`
	Email    string    `json:"email"`
	UserType string    `json:"userType"`
	Password string    `json:"password"`
}

const (
	Admin   = "AdminUser"
	Regular = "RegularUser"
)

func UnMarshalUser(request io.ReadCloser) *User {
	user := User{}
	body, err := io.ReadAll(request)
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(body, &user)
	if err != nil {
		log.Fatal(err)
	}
	return &user
}
