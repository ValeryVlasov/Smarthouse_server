package service

import (
	"github.com/ValeryVlasov/Smarthouse_server"
	"github.com/ValeryVlasov/Smarthouse_server/pkg/repository"
)

type DeviceDetectorService struct {
	repo repository.DeviceDetector
}

func NewDeviceDetectorService(repo repository.DeviceDetector) *DeviceDetectorService {
	return &DeviceDetectorService{repo: repo}
}

func (s *DeviceDetectorService) Create(userId int, detector Smarthouse_server.DeviceDetector) (int, error) {
	return s.repo.Create(userId, detector)
}

func (s *DeviceDetectorService) GetAll(userId int) ([]Smarthouse_server.DeviceDetector, error) {
	return s.repo.GetAll(userId)
}

func (s *DeviceDetectorService) GetById(userId, detectorId int) (Smarthouse_server.DeviceDetector, error) {
	return s.repo.GetById(userId, detectorId)
}

func (s *DeviceDetectorService) Delete(userId, detectorId int) error {
	return s.repo.Delete(userId, detectorId)
}

func (s *DeviceDetectorService) Update(userId, detectorId int, input Smarthouse_server.UpdateDetectorInput) error {
	if err := input.Validate(); err != nil {
		return err
	}
	return s.repo.Update(userId, detectorId, input)
}
