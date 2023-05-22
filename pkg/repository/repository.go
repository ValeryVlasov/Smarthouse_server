package repository

import (
	"github.com/ValeryVlasov/Smarthouse_server"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user Smarthouse_server.User) (int, error)
	GetUser(username, password string) (Smarthouse_server.User, error)
	GetUser2(username, password string) (Smarthouse_server.User2, error)
	IsSameUser(login, password interface{}) bool
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

type Repository struct {
	Authorization
	DeviceLight
	DeviceCamera
	DeviceDetector
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization:  NewAuthPostgres(db),
		DeviceLight:    NewDeviceLightPostgres(db),
		DeviceCamera:   NewDeviceCameraPostgres(db),
		DeviceDetector: NewDeviceDetectorPostgres(db),
	}
}
