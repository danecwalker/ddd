package client

import "github.com/oklog/ulid/v2"

type Client struct {
	id          ulid.ULID
	advisorID   ulid.ULID
	name        string
	email       string
	phoneNumber string
}

func NewClient(advisorID ulid.ULID, name, email, phoneNumber string) Client {
	return Client{
		id:          ulid.MustNew(ulid.Now(), nil),
		advisorID:   advisorID,
		name:        name,
		email:       email,
		phoneNumber: phoneNumber,
	}
}

func (c *Client) ID() ulid.ULID {
	return c.id
}

func (c *Client) SetID(id ulid.ULID) {
	c.id = id
}

func (c *Client) AdvisorID() ulid.ULID {
	return c.advisorID
}

func (c *Client) Name() string {
	return c.name
}

func (c *Client) Email() string {
	return c.email
}

func (c *Client) PhoneNumber() string {
	return c.phoneNumber
}
