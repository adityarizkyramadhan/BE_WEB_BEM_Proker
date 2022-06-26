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
	if err := s.db.Model(&domain.EntitasProker{}).Preload("SaveImages").Find(&datas).Error; err != nil {
		return []domain.EntitasProker{}, err
	}
	return datas, nil
}

func (s serviceProker) GetByID(id uint) (domain.EntitasProker, error) {
	var data domain.EntitasProker
	if err := s.db.Model(&domain.EntitasProker{}).Where("id = ?", id).Preload("SaveImages").Take(&data).Error; err != nil {
		return domain.EntitasProker{}, err
	}
	return data, nil
}

func (s serviceProker) Create(proker *domain.EntitasProker) error {
	if err := s.db.Model(&domain.EntitasProker{}).Create(&proker).Error; err != nil {
		return err
	}
	return nil
}

func (s serviceProker) Update(id uint, proker *domain.EntitasProker) error {
	if err := s.db.Model(&domain.EntitasProker{}).Where("id = ?", id).Save(&proker).Error; err != nil {
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

func (s serviceProker) Login(email string, password string) (*domain.Admin, error) {
	var data domain.Admin
	if err := s.db.Model(&domain.Admin{}).Where("email = ? AND password = ?", email, password).Take(&data).Error; err != nil {
		return &domain.Admin{}, err
	}
	return &data, nil
}

func (s serviceProker) Register(admin *domain.Admin) (*domain.Admin, error) {
	if err := s.db.Model(&domain.Admin{}).Create(&admin).Error; err != nil {
		return &domain.Admin{}, err
	}
	return admin, nil
}
