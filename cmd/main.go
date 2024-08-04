package main

import (
	"fmt"

	"github.com/danecwalker/progo/internal/app"
	"github.com/danecwalker/progo/internal/app/commands/account"
	"github.com/danecwalker/progo/internal/app/commands/advisor"
	"github.com/danecwalker/progo/internal/infrastructure/api"
	"github.com/danecwalker/progo/internal/infrastructure/db/sqlite"
	"github.com/danecwalker/progo/pkg/cqrs"
)

func init() {
	db := sqlite.NewDB()

	clientRepo := sqlite.NewClientRepository(db)

	cqrs.RegisterHandler[account.CreateAccountCommand, string](account.NewCreateAccountCommandHandler())
	cqrs.RegisterHandler[advisor.AddClientCommand, string](advisor.NewAddClientCommandHandler(clientRepo))
}

func main() {
	apiMux := api.NewApiHandler()

	app.RegisterAccountRoutes(apiMux)

	if err := apiMux.Serve(":4356"); err != nil {
		fmt.Println(err)
	}
}
