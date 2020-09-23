package db_users

import (
	"database/sql"
	"github.com/go-sql-driver/mysql"
	"log"
)

var (
	Client *sql.DB
)

func init() {
	dataSourceName := "root:@tcp(localhost:3306)/db_users?charset=utf8"

	var err error
	Client, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		panic(err)
	}

	if err = Client.Ping(); err != nil {
		panic(err)
	}

	mysql.SetLogger(nil)

	log.Println("Database configured successfully")
}
