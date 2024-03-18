package views

import "github.com/google/uuid"

type UserResponse struct {
	UID      uuid.UUID `json:"uid"`
	Name     string    `json:"name"`
	Email    string    `json:"email"`
	UserType string    `json:"userType"`
	Token    string    `json:"token"`
}
