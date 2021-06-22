package users

import (
	"fmt"
	"sample-go-gin-api/datasource/mysql/users_db"
	"sample-go-gin-api/utils/errors"
)

const (
	queryInsertUser = "INSERT INTO users(first_name, last_name, email, password) VALUES(?, ?, ?, ?);"
	queryGetUser    = "SELECT id, first_name, last_name, email FROM users WHERE id = ?;"
)

func (user *User) Get() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryGetUser)
	if err != nil {
		return errors.NewInternalServerError("database error")
	}
	defer stmt.Close()

	result := stmt.QueryRow(user.Id)
	if getErr := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email); getErr != nil {
		return errors.NewInternalServerError("database error")
	}

	return nil
}

func (user *User) Save() *errors.RestErr {
	fmt.Println("Save")
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
	if err != nil {
		return errors.NewInternalServerError("database error")
	}

	user.Id = userId

	return nil
}
