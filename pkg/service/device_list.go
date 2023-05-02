package service

import (
	"github.com/ValeryVlasov/Smarthouse_server"
	"github.com/ValeryVlasov/Smarthouse_server/pkg/repository"
)

type DeviceListService struct {
	repo repository.DeviceList
}

func NewDeviceListService(repo repository.DeviceList) *DeviceListService {
	return &DeviceListService{repo: repo}
}

func (s *DeviceListService) Create(userId int, list Smarthouse_server.DeviceList) (int, error) {
	return s.repo.Create(userId, list)
}

func (s *DeviceListService) GetAll(userId int) ([]Smarthouse_server.DeviceList, error) {
	return s.repo.GetAll(userId)
}

func (s *DeviceListService) GetById(userId, listId int) (Smarthouse_server.DeviceList, error) {
	return s.repo.GetById(userId, listId)
}

func (s *DeviceListService) Delete(userId, listId int) error {
	return s.repo.Delete(userId, listId)
}

func (s *DeviceListService) Update(userId, listId int, input Smarthouse_server.UpdateListInput) error {
	if err := input.Validate(); err != nil {
		return err
	}
	return s.repo.Update(userId, listId, input)
}
