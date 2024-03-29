package users

import (
	"fmt"
	"log"
	"sample-go-gin-api/datasource/mysql/users_db"
	"sample-go-gin-api/utils/errors"

	_ "github.com/go-sql-driver/mysql"
)

const (
	queryInsertUser   = "INSERT INTO users(first_name, last_name, email, password, status) VALUES (?, ?, ?, ?, ?)"
	queryGetUser      = "SELECT id, first_name, last_name, email FROM users WHERE id = ?"
	queryUpdateUser   = "UPDATE users SET first_name = ?, last_name = ?, email = ? WHERE id = ?"
	queryDeleteUser   = "DELETE FROM users WHERE id = ?"
	queyrFindByStatus = "SELECT id, first_name, last_name, email, status FROM users WHERE status = ?"
)

func (user *User) Get() *errors.RestErr {
	if users_db.Client == nil {
		return errors.NewInternalServerError("users_db.Client error")
	}
	log.Println(users_db.Client)

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

	result, saveErr := stmt.Exec(user.FirstName, user.LastName, user.Email, user.Password, user.Status)
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

func (user *User) Update() *errors.RestErr {
	fmt.Println("Update")
	stmt, err := users_db.Client.Prepare(queryUpdateUser)
	if err != nil {
		return errors.NewInternalServerError("database error")
	}
	defer stmt.Close()

	fmt.Println("Exec")
	_, updateErr := stmt.Exec(user.FirstName, user.LastName, user.Email, user.Id)
	if updateErr != nil {
		return errors.NewInternalServerError("database error")
	}

	return nil
}

func (user *User) Delete() *errors.RestErr {
	fmt.Println("Delete")
	stmt, err := users_db.Client.Prepare(queryDeleteUser)
	if err != nil {
		return errors.NewInternalServerError("database error")
	}
	defer stmt.Close()

	if _, err := stmt.Exec(user.Id); err != nil {
		return errors.NewInternalServerError("database error")
	}

	return nil
}

func (user *User) FindByStatus(status string) ([]User, *errors.RestErr) {
	stmt, err := users_db.Client.Prepare(queyrFindByStatus)
	if err != nil {
		return nil, errors.NewInternalServerError("database error")
	}
	defer stmt.Close()

	rows, err := stmt.Query(status)
	if err != nil {
		return nil, errors.NewInternalServerError("database error")
	}
	defer rows.Close()

	results := make([]User, 0)
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.Status); err != nil {
			return nil, errors.NewInternalServerError("database error")
		}
		results = append(results, user)
	}

	if len(results) == 0 {
		return nil, errors.NewInternalServerError(fmt.Sprintf("ステータスが一致するユーザが存在しませんでした。 %s", status))
	}

	return results, nil

}
