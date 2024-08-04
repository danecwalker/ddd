package app

import (
	"github.com/danecwalker/progo/internal/app/commands/account"
	"github.com/danecwalker/progo/internal/infrastructure/api"
)

func RegisterAccountRoutes(_api *api.ApiHandler) {
	subRouter := _api.NewSubRouter("/account")

	subRouter.HandleFunc("POST /", account.CreateAccount)
}
