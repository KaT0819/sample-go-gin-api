package users

import (
	"sample-go-gin-api/utils/errors"
	"strings"
)

type User struct {
	Id        int64
	FirstName string
	LastName  string
	Email     string
	Password  string
}

func (user *User) Validate() *errors.RestErr {
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	if user.Email == "" {
		return errors.NewBadRequestError("不正なメールアドレスです")
	}

	user.Password = strings.TrimSpace(user.Password)
	if user.Password == "" {
		return errors.NewBadRequestError("不正なパスワードです")
	}

	return nil
}
