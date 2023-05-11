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

type DeviceLight interface {
	Create(userId int, light Smarthouse_server.DeviceLight) (int, error)
	GetAll(userId int) ([]Smarthouse_server.DeviceLight, error)
	GetById(userId, lightId int) (Smarthouse_server.DeviceLight, error)
	Delete(userId, lightId int) error
	Update(userId, lightId int, input Smarthouse_server.UpdateLightInput) error
}

type DeviceCamera interface {
	Create(userId int, camera Smarthouse_server.DeviceCamera) (int, error)
	GetAll(userId int) ([]Smarthouse_server.DeviceCamera, error)
	GetById(userId, cameraId int) (Smarthouse_server.DeviceCamera, error)
	Delete(userId, cameraId int) error
	Update(userId, cameraId int, input Smarthouse_server.UpdateCameraInput) error
}

type DeviceDetector interface {
	Create(userId int, detector Smarthouse_server.DeviceDetector) (int, error)
	GetAll(userId int) ([]Smarthouse_server.DeviceDetector, error)
	GetById(userId, detectorId int) (Smarthouse_server.DeviceDetector, error)
	Delete(userId, detectorId int) error
	Update(userId, detectorId int, input Smarthouse_server.UpdateDetectorInput) error
}

type Service struct {
	Authorization
	DeviceList
	DeviceItem
	DeviceLight
	DeviceCamera
	DeviceDetector
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization:  NewAuthService(repos.Authorization),
		DeviceList:     NewDeviceListService(repos.DeviceList),
		DeviceItem:     NewDeviceItemService(repos.DeviceItem, repos.DeviceList),
		DeviceLight:    NewDeviceLightService(repos.DeviceLight),
		DeviceCamera:   NewDeviceCameraService(repos.DeviceCamera),
		DeviceDetector: NewDeviceDetectorService(repos.DeviceDetector),
	}
}
