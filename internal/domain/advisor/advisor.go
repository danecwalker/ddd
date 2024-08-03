package advisor

import (
	"github.com/oklog/ulid/v2"
)

type Advisor struct {
	id    ulid.ULID
	name  string
	email string
}

func NewAdvisor(name, email string) *Advisor {
	return &Advisor{
		id:    ulid.MustNew(ulid.Now(), nil),
		name:  name,
		email: email,
	}
}

func (a *Advisor) ID() ulid.ULID {
	return a.id
}
