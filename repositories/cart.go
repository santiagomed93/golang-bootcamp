package repositories

import "database/sql"

type DatabaseInterface interface {
	Query(query string) (sql.Result, error)
}

func createCart(db DatabaseInterface, items []Items) {
	db.Query
}
