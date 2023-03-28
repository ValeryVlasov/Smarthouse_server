package service

import (
	"github.com/ValeryVlasov/Smarthouse_server"
	"github.com/ValeryVlasov/Smarthouse_server/pkg/repository"
)

type Authorization interface {
	CreateUser(user Smarthouse_server.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type Service struct {
	Authorization
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}
