package database

import (
	"database/sql"
	"log"
)

var DbConn *sql.DB

func SetupDatabase() *sql.DB {
	var err error
	DbConn, err = sql.Open("mysql", "root:MySQL123@tcp(127.0.0.1:3306)/shoppingDB")
	if err != nil {
		log.Fatal(err)
	}
	return DbConn
}
