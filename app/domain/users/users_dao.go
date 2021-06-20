package users

import (
	"sample-go-gin-api/app/datasource/mysql/users_db"
	"sample-go-gin-api/utils/errors"
)

var (
	queryInsertUser = "INSERT INTO users(first_name, last_name, email, password VALUES(?, ?, ?, ?)"
)

func (user *User) Save() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryInsertUser)
	if err != nil {
		return errors.NewInternalServerError("database error")
	}
	defer stmt.Close()

	result, saveErr := stmt.Exec(user.FirstName, user.LastName, user.Email, user.Password)
	if saveErr != nil {
		return errors.NewInternalServerError("database error")
	}

	userId, err := result.LastInsertId()
	if saveErr != nil {
		return errors.NewInternalServerError("database error")
	}

	user.Id = userId

	return nil
}
