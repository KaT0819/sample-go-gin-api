package services

import (
	"fmt"
	"sample-go-gin-api/domain/users"
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

func GetUser(userId int64) (*users.User, *errors.RestErr) {
	result := &users.User{Id: userId}
	if err := result.Get(); err != nil {
		return nil, err
	}
	return result, nil
}

func UpdateUser(user users.User, isPertial bool) (*users.User, *errors.RestErr) {
	fmt.Println("UpdateUser start")
	current, err := GetUser(user.Id)
	if err != nil {
		return nil, err
	}

	if isPertial {
		if user.FirstName != "" {
			current.FirstName = user.FirstName
		}
		if user.LastName != "" {
			current.LastName = user.LastName
		}
		if user.Email != "" {
			current.Email = user.Email
		}
	} else {
		current.FirstName = user.FirstName
		current.LastName = user.LastName
		current.Email = user.Email
	}

	// update
	if err := current.Update(); err != nil {
		return nil, err
	}

	return current, nil
}

func DeleteUser(userId int64) *errors.RestErr {
	user := &users.User{Id: userId}

	// delete
	if err := user.Delete(); err != nil {
		return err
	}

	return nil
}
