package auth

import (
	"time"
)

type Core struct {
	ID        int
	Username  string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Business interface {
	Login(data Core) (token string, ID int, err error)
	Register(data Core) error
}

type Data interface {
	FindUser(email string) (Core, error)
	InsertUser(data Core) error
}
