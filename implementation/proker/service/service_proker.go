package service

import (
	"BE_WEB_BEM_Proker/domain"

	"gorm.io/gorm"
)

type serviceProker struct {
	db *gorm.DB
}

func NewServiceProker(db *gorm.DB) domain.ProkerService {
	return &serviceProker{db}
}

func (s serviceProker) GetAll() ([]domain.EntitasProker, error) {
	var datas []domain.EntitasProker
	if err := s.db.Model(domain.EntitasProker{}).Find(&datas).Error; err != nil {
		return []domain.EntitasProker{}, err
	}
	return datas, nil
}

func (s serviceProker) GetByID(id uint) (domain.EntitasProker, error) {
	var data domain.EntitasProker
	if err := s.db.Model(domain.EntitasProker{}).Where("id = ?", id).Take(&data).Error; err != nil {
		return domain.EntitasProker{}, err
	}
	return data, nil
}

func (s serviceProker) Create(proker *domain.EntitasProker) error {
	if err := s.db.Model(domain.EntitasProker{}).Create(&proker).Error; err != nil {
		return err
	}
	return nil
}

func (s serviceProker) Update(id uint, proker *domain.EntitasProker) error {
	if err := s.db.Model(domain.EntitasProker{}).Where("id = ?", id).Save(&proker).Error; err != nil {
		return err
	}
	return nil
}

func (s serviceProker) Delete(id uint) error {
	if err := s.db.Delete("id = ?", id).Error; err != nil {
		return err
	}
	return nil
}