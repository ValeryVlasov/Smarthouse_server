package service

import (
	"github.com/ValeryVlasov/Smarthouse_server"
	"github.com/ValeryVlasov/Smarthouse_server/pkg/repository"
)

type Authorization interface {
	CreateUser(user Smarthouse_server.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
	IsSameUser(login, password interface{}) (Smarthouse_server.User2, bool)
}

type DeviceLight interface {
	Create(userId int, light Smarthouse_server.DeviceLight) (int, error)
	GetAll(userId int) ([]Smarthouse_server.DeviceLight, error)
	GetById(userId, lightId int) (Smarthouse_server.DeviceLight, error)
	Delete(userId, lightId int) error
	Toggle(userId, lightId int) error
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
	DeviceLight
	DeviceCamera
	DeviceDetector
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization:  NewAuthService(repos.Authorization),
		DeviceLight:    NewDeviceLightService(repos.DeviceLight),
		DeviceCamera:   NewDeviceCameraService(repos.DeviceCamera),
		DeviceDetector: NewDeviceDetectorService(repos.DeviceDetector),
	}
}
