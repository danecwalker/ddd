package sqlite

import (
	"context"
	"database/sql"
	"fmt"
)

var _db *sql.DB = nil

func NewDB() *sql.DB {
	db, err := sql.Open("sqlite3", "./progo.db")
	if err != nil {
		panic(err)
	}

	_db = db

	return db
}

func WithUnitOfWork(ctx context.Context, fn func(_dbContext context.Context) error) error {
	if _db == nil {
		NewDB()
	}

	tx, err := _db.Begin()
	if err != nil {
		return err
	}

	_dbtx := context.WithValue(ctx, txKey{}, tx)

	if err := fn(_dbtx); err != nil {
		fmt.Println("Rolling back transaction")
		err = tx.Rollback()
		return err
	}

	return tx.Commit()
}
