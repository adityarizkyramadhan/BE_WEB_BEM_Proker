package service

import (
	"BE_WEB_BEM_Proker/config"
	"BE_WEB_BEM_Proker/model"
)

var DB, _ = config.InitDB()

type ServiceBPH interface {
	Create(*model.EntitasProker) (*model.EntitasProker, error)
	Read() (*[]model.EntitasProker, error)
	Update(*model.EntitasProker) (*model.EntitasProker, error)
	Delete(*model.EntitasProker) (*model.EntitasProker, error)
}

type bPHService struct{}

func NewBPHService() ServiceBPH {
	return &bPHService{}
}

func (s *bPHService) Create(entitas *model.EntitasProker) (*model.EntitasProker, error) {
	if err := DB.Create(&entitas).Error; err != nil {
		return nil, err
	}
	return entitas, nil
}

func (s *bPHService) Update(entitas *model.EntitasProker) (*model.EntitasProker, error) {
	if err := DB.Save(&entitas).Error; err != nil {
		return nil, err
	}
	return entitas, nil
}

func (s *bPHService) Delete(entitas *model.EntitasProker) (*model.EntitasProker, error) {
	if err := DB.Delete(&entitas).Error; err != nil {
		return nil, err
	}
	return entitas, nil
}

func (s *bPHService) Read() (*[]model.EntitasProker, error) {
	var entitasBanyak *[]model.EntitasProker
	if err := DB.Find(&entitasBanyak).Error; err != nil {
		return nil, err
	}
	return entitasBanyak, nil
}
