package sqlite

import (
	"context"
	"database/sql"

	"github.com/danecwalker/progo/internal/domain/client"
	_ "github.com/mattn/go-sqlite3"
)

type ClientRepository struct {
	db *sql.DB
}

func NewClientRepository(ctx *sql.DB) *ClientRepository {
	return &ClientRepository{db: ctx}
}

func (s *ClientRepository) Save(ctx context.Context, client client.Client) error {
	store := getDBTX(ctx, s.db)

	_, err := store.Exec("INSERT INTO clients (ID, advisorID, name, email, phoneNumber) VALUES (?, ?, ?, ?, ?)", client.ID().String(), client.AdvisorID().String(), client.Name(), client.Email(), client.PhoneNumber())
	if err != nil {
		return err
	}
	return nil
}
