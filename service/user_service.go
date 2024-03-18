package service

import (
	log "github.com/sirupsen/logrus"
	"gomvc/resources"
	"gomvc/util"
	"gomvc/views"
)

type UserService struct {
}

func NewUserService() *UserService {
	return &UserService{}
}

func (us *UserService) LoginUser(email, password string) (views.UserResponse, error) {
	res, err := resources.LoginUser(email, password)

	if err != nil {
		log.Error(err)
		return views.UserResponse{}, err
	}

	res.Token = util.GenerateNewTokenForUser(res.Email, res.UserType)

	return *res, nil
}

func (us *UserService) ValidateToken(tokenString string) error {
	return util.VerifyToken(tokenString[7:])
}
