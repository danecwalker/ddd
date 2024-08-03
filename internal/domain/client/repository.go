package client

import "context"

type IClientRepository interface {
	Save(context context.Context, client Client) error
}
