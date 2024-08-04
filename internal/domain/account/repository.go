package account

import "github.com/oklog/ulid/v2"

type IAccountRepository interface {
	Save(account Account) error
	FindByEmail(email string) (Account, error)
	FindByID(id ulid.ULID) (Account, error)
}
