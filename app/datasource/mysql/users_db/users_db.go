package users_db

import (
	"database/sql"
	"fmt"
	"log"
)

var (
	Client   *sql.DB
	username = "docker"
	password = "docker"
	host     = "127.0.0.1:3306"
	schema   = "database"
)

func init() {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8", username, password, host, schema)

	Client, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		panic(err)
	}

	if err = Client.Ping(); err != nil {
		panic(err)
	}

	log.Println("database successfully configured")
}
