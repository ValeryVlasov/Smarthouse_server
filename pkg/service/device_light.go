package service

import (
	"github.com/ValeryVlasov/Smarthouse_server"
	"github.com/ValeryVlasov/Smarthouse_server/pkg/repository"
)

type DeviceLightService struct {
	repo repository.DeviceLight
}

func NewDeviceLightService(repo repository.DeviceLight) *DeviceLightService {
	return &DeviceLightService{repo: repo}
}

func (s *DeviceLightService) Create(userId int, light Smarthouse_server.DeviceLight) (int, error) {
	return s.repo.Create(userId, light)
}

func (s *DeviceLightService) GetAll(userId int) ([]Smarthouse_server.DeviceLight, error) {
	return s.repo.GetAll(userId)
}

func (s *DeviceLightService) GetById(userId, lightId int) (Smarthouse_server.DeviceLight, error) {
	return s.repo.GetById(userId, lightId)
}

func (s *DeviceLightService) Delete(userId, lightId int) error {
	return s.repo.Delete(userId, lightId)
}

func (s *DeviceLightService) Toggle(userId, lightId int) error {
	return s.repo.Toggle(userId, lightId)
}

func (s *DeviceLightService) Update(userId, lightId int, input Smarthouse_server.UpdateLightInput) error {
	if err := input.Validate(); err != nil {
		return err
	}
	return s.repo.Update(userId, lightId, input)
}
