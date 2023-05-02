package repository

import (
	"github.com/ValeryVlasov/Smarthouse_server"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user Smarthouse_server.User) (int, error)
	GetUser(username, password string) (Smarthouse_server.User, error)
}

type DeviceList interface {
	Create(userId int, list Smarthouse_server.DeviceList) (int, error)
	GetAll(userId int) ([]Smarthouse_server.DeviceList, error)
	GetById(userId, listId int) (Smarthouse_server.DeviceList, error)
	Delete(userId, listId int) error
	Update(userId, listId int, input Smarthouse_server.UpdateListInput) error
}

type DeviceItem interface {
	Create(listId int, item Smarthouse_server.DeviceItem) (int, error)
	GetAll(userId, listId int) ([]Smarthouse_server.DeviceItem, error)
	GetById(userId, itemId int) (Smarthouse_server.DeviceItem, error)
	Delete(userId, itemId int) error
	Update(userId, itemId int, input Smarthouse_server.UpdateItemInput) error
}

type Repository struct {
	Authorization
	DeviceList
	DeviceItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		DeviceList:    NewDeviceListPostgres(db),
		DeviceItem:    NewDeviceItemPostgres(db),
	}
}
