package users_db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
)

const (
	//You must configure it yourself at your environment variable
	MYSQL_USERNAME = "MYSQL_USERNAME"
	MYSQL_PASSWORD = "MYSQL_PASSWORD"
	MYSQL_HOST     = "MYSQL_HOST"
	MYSQL_DB       = "MYSQL_DB"
)

var (
	Client   *sql.DB
	username = os.Getenv(MYSQL_USERNAME)
	password = os.Getenv(MYSQL_PASSWORD)
	host     = os.Getenv(MYSQL_HOST)
	schema   = os.Getenv(MYSQL_DB)
)

func init() {
	datasourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s",
		username,
		password,
		host,
		schema,
	)

	var err error
	Client, err = sql.Open("mysql", datasourceName)
	if err != nil {
		panic(err)
	}
	if err := Client.Ping(); err != nil {
		panic(err)
	}
	log.Println("Database successfully configured")
}
