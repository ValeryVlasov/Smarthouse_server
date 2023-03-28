package repository

import (
	"github.com/ValeryVlasov/Smarthouse_server"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user Smarthouse_server.User) (int, error)
	GetUser(username, password string) (Smarthouse_server.User, error)
}

type Repository struct {
	Authorization
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}
