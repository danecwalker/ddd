package account

import (
	"context"
	"net/http"

	"github.com/danecwalker/progo/internal/domain/account"
	"github.com/danecwalker/progo/internal/infrastructure/api"
	"github.com/danecwalker/progo/pkg/cqrs"
)

type CreateAccountCommand struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type CreateAccountCommandHandler struct {
	accountRepo account.IAccountRepository
}

func NewCreateAccountCommandHandler() CreateAccountCommandHandler {
	return CreateAccountCommandHandler{
		accountRepo: nil,
	}
}

func (h CreateAccountCommandHandler) Handle(command CreateAccountCommand, ctx context.Context) (string, error) {
	account := account.NewAccount(command.Email, command.Password)

	// if err := h.accountRepo.Save(account); err != nil {
	// 	return "", err
	// }

	return account.ID().String(), nil
}

func CreateAccount(res http.ResponseWriter, req *http.Request) {
	// Create account
	command, err := api.CommandFromBody[CreateAccountCommand](req)
	if err != nil {
		api.ErrorResponse(res, 500, err.Error())
		return
	}

	// Handle command
	result := cqrs.Dispatch[CreateAccountCommand, string](command, context.Background())
	if result.IsError() {
		api.ErrorResponse(res, 500, result.Error())
		return
	}

	// Return account ID
	api.JSONResponse(res, 200, api.Json{
		"account_id": result.Value,
	})
}
