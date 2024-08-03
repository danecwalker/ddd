package advisor

import (
	"bufio"
	"context"
	"fmt"
	"os"

	"github.com/danecwalker/progo/internal/domain/client"
	"github.com/danecwalker/progo/internal/infrastructure/db/sqlite"
	"github.com/oklog/ulid/v2"
)

type AddClientCommand struct {
	AdvisorID   string `json:"advisor_id"`
	ClientName  string `json:"client_name"`
	ClientEmail string `json:"client_email"`
	ClientPhone string `json:"client_phone"`
}

type AddClientCommandHandler struct {
	clientRepo client.IClientRepository
}

func NewAddClientCommandHandler(clientRepo client.IClientRepository) AddClientCommandHandler {
	return AddClientCommandHandler{
		clientRepo: clientRepo,
	}
}

func (h AddClientCommandHandler) Handle(command AddClientCommand, ctx context.Context) (string, error) {
	client := client.NewClient(ulid.MustParse(command.AdvisorID), command.ClientName, command.ClientEmail, command.ClientPhone)

	err := sqlite.WithUnitOfWork(ctx, func(_dbContext context.Context) error {
		// Save client to database
		if err := h.clientRepo.Save(ctx, client); err != nil {
			return err
		}

		client.SetID(ulid.MustParse(command.AdvisorID))
		// Save client to database
		if err := h.clientRepo.Save(ctx, client); err != nil {
			return err
		}

		// wait for user to confirm
		fmt.Println("Continue...")
		_, _ = bufio.NewReader(os.Stdin).ReadBytes('\n')

		return nil
	})

	if err != nil {
		return "", err
	}

	return command.AdvisorID, nil
}
