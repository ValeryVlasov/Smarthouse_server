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

type DeviceList interface {
	Create(userId int, list Smarthouse_server.DeviceList) (int, error)
	GetAll(userId int) ([]Smarthouse_server.DeviceList, error)
	GetById(userId, listId int) (Smarthouse_server.DeviceList, error)
	Delete(userId, listId int) error
	Update(userId, listId int, input Smarthouse_server.UpdateListInput) error
}

type DeviceItem interface {
	Create(userId, listId int, item Smarthouse_server.DeviceItem) (int, error)
	GetAll(userId, listId int) ([]Smarthouse_server.DeviceItem, error)
	GetById(userId, itemId int) (Smarthouse_server.DeviceItem, error)
	Delete(userId, itemId int) error
	Update(userId, itemId int, input Smarthouse_server.UpdateItemInput) error
}

type Service struct {
	Authorization
	DeviceList
	DeviceItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		DeviceList:    NewDeviceListService(repos.DeviceList),
		DeviceItem:    NewDeviceItemService(repos.DeviceItem, repos.DeviceList),
	}
}
