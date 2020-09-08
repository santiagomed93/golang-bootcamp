package database

import (
	"database/sql"
	"log"
)

type Database struct {
	Conn *sql.DB
}

func SetupDatabase() *sql.DB {
	DbConn, err := sql.Open("mysql", "root:MySQL123@tcp(127.0.0.1:3306)/shoppingDB")
	if err != nil {
		log.Fatal(err)
	}
	return DbConn
}

func NewDatabase() *Database {
	return &Database{
		Conn: SetupDatabase(),
	}
}
