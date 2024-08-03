package main

import (
	"context"
	"fmt"

	"github.com/danecwalker/progo/internal/app/commands/advisor"
	"github.com/danecwalker/progo/internal/infrastructure/db/sqlite"
	"github.com/danecwalker/progo/pkg/cqrs"
)

func init() {
	db := sqlite.NewDB()

	clientRepo := sqlite.NewClientRepository(db)

	cqrs.RegisterHandler[advisor.AddClientCommand, string](advisor.NewAddClientCommandHandler(clientRepo))
}

func main() {
	ctx, _ := context.WithCancel(context.Background())
	v := cqrs.Dispatch[advisor.AddClientCommand, string](advisor.AddClientCommand{
		AdvisorID:   "01F0ZQZQZQZQZQZQZQZQZQZQZQ",
		ClientName:  "John Doe",
		ClientEmail: "johndoe@mail.com",
		ClientPhone: "555-555-5555",
	}, ctx)

	if v.IsError() {
		fmt.Println(v.Err.Error())
	}

	fmt.Println(v)
}
