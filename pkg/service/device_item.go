package service

import (
	"github.com/ValeryVlasov/Smarthouse_server"
	"github.com/ValeryVlasov/Smarthouse_server/pkg/repository"
)

type DeviceItemService struct {
	repo     repository.DeviceItem
	listRepo repository.DeviceList
}

func NewDeviceItemService(repo repository.DeviceItem, listRepo repository.DeviceList) *DeviceItemService {
	return &DeviceItemService{repo: repo, listRepo: listRepo}
}

func (s *DeviceItemService) Create(userId, listId int, item Smarthouse_server.DeviceItem) (int, error) {
	_, err := s.listRepo.GetById(userId, listId)
	if err != nil {
		return 0, err
	}

	return s.repo.Create(listId, item)
}

func (s *DeviceItemService) GetAll(userId, listId int) ([]Smarthouse_server.DeviceItem, error) {
	return s.repo.GetAll(userId, listId)
}

func (s *DeviceItemService) GetById(userId, itemId int) (Smarthouse_server.DeviceItem, error) {
	return s.repo.GetById(userId, itemId)
}

func (s *DeviceItemService) Delete(userId, itemId int) error {
	return s.repo.Delete(userId, itemId)
}

func (s *DeviceItemService) Update(userId, itemId int, input Smarthouse_server.UpdateItemInput) error {
	if err := input.Validate(); err != nil {
		return err
	}
	return s.repo.Update(userId, itemId, input)
}
