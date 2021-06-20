package services

import (
	"sample-go-gin-api/app/domain/users"
	"sample-go-gin-api/utils/errors"
)

func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	// validate
	if err := user.Validate(); err != nil {
		return nil, err
	}
	// save
	if err := user.Save(); err != nil {
		return nil, err
	}

	return &user, nil
}
