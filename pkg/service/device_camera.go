package service

import (
	"github.com/ValeryVlasov/Smarthouse_server"
	"github.com/ValeryVlasov/Smarthouse_server/pkg/repository"
)

type DeviceCameraService struct {
	repo repository.DeviceCamera
}

func NewDeviceCameraService(repo repository.DeviceCamera) *DeviceCameraService {
	return &DeviceCameraService{repo: repo}
}

func (s *DeviceCameraService) Create(userId int, camera Smarthouse_server.DeviceCamera) (int, error) {
	return s.repo.Create(userId, camera)
}

func (s *DeviceCameraService) GetAll(userId int) ([]Smarthouse_server.DeviceCamera, error) {
	return s.repo.GetAll(userId)
}

func (s *DeviceCameraService) GetById(userId, cameraId int) (Smarthouse_server.DeviceCamera, error) {
	return s.repo.GetById(userId, cameraId)
}

func (s *DeviceCameraService) Delete(userId, cameraId int) error {
	return s.repo.Delete(userId, cameraId)
}

func (s *DeviceCameraService) Update(userId, cameraId int, input Smarthouse_server.UpdateCameraInput) error {
	if err := input.Validate(); err != nil {
		return err
	}
	return s.repo.Update(userId, cameraId, input)
}
