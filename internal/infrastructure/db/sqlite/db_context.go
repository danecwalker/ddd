package sqlite

import (
	"context"
	"database/sql"
	"fmt"
)

type DBorTX interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
	Query(query string, args ...interface{}) (*sql.Rows, error)
	QueryRow(query string, args ...interface{}) *sql.Row
}

type txKey struct{}

func getDBTX(ctx context.Context, _db *sql.DB) DBorTX {
	// returns the transaction from the context
	// if it doesn't exist, it will return _db
	if tx, ok := ctx.Value(txKey{}).(*sql.Tx); ok {
		fmt.Println("Context has transaction")
		return tx
	} else {
		fmt.Println("Context has no transaction")
		return _db
	}
}
