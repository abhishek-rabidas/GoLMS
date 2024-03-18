package resources

import (
	"github.com/google/uuid"
	"gomvc/exception"
	"gomvc/model"
	"gomvc/util"
	"gomvc/views"
)

var users = map[string]model.User{}

func PopulateUsers() {

	password, err := util.HashPassword("Test@121")

	if err != nil {
		panic(err)
		return
	}

	user1 := model.User{
		UID:      uuid.New(),
		Name:     "Abhishek",
		Email:    "abhishekrabidas07@gmail.com",
		UserType: model.Admin,
		Password: password,
	}

	user2 := model.User{
		UID:      uuid.New(),
		Name:     "Raj",
		Email:    "abhishekrabidas7@gmail.com",
		UserType: model.Regular,
		Password: password,
	}

	users["abhishekrabidas07@gmail.com"] = user1
	users["abhishekrabidas7@gmail.com"] = user2
	
}

func LoginUser(email, password string) (*views.UserResponse, error) {
	user, isExists := users[email]

	if !isExists {
		err := exception.New("Unable retrieve user")
		return nil, err
	}

	if util.CheckPasswordHash(password, user.Password) {
		response := views.UserResponse{
			UID:      user.UID,
			Name:     user.Name,
			Email:    user.Email,
			UserType: user.UserType,
			Token:    "",
		}
		return &response, nil
	} else {
		err := exception.New("Failed to authenticate")
		return nil, err
	}
}
