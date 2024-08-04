package account

import (
	"github.com/oklog/ulid/v2"
)

type Account struct {
	id        ulid.ULID
	advisorId *ulid.ULID
	email     string
	password  string
}

func NewAccount(email, password string) Account {
	return Account{
		id:        ulid.MustNew(ulid.Now(), nil),
		advisorId: nil,
		email:     email,
		password:  password,
	}
}

func (a *Account) SetAdvisorID(advisorID ulid.ULID) {
	a.advisorId = &advisorID
}

func (a Account) ID() ulid.ULID {
	return a.id
}

func (a Account) AdvisorID() *ulid.ULID {
	return a.advisorId
}

func (a Account) Email() string {
	return a.email
}

func (a Account) Password() string {
	return a.password
}
