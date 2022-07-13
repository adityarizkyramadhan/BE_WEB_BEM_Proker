package db

import (
	"BE_WEB_BEM_Proker/domain"

	"gorm.io/gorm"
)

type prokerDatabase struct {
	db *gorm.DB
}

type DatabaseService interface {
	GetAll() (*[]domain.EntitasProker, error)
	CreateProker(proker domain.EntitasProker) error
	GetByID(id uint) (*domain.EntitasProker, error)
	Delete(id uint) error
}

var prokerDB *prokerDatabase

//Biar db yang dikirim datanya cuma itu itu saja
func InitProkerDB(db *gorm.DB) DatabaseService {
	if prokerDB == nil {
		prokerDB = &prokerDatabase{db}
	}
	return prokerDB
}

func (p *prokerDatabase) GetAll() (*[]domain.EntitasProker, error) {
	var prokers []domain.EntitasProker
	err := p.db.Find(&prokers).Error
	if err != nil {
		return nil, err
	}
	return &prokers, nil
}

func (p *prokerDatabase) CreateProker(proker domain.EntitasProker) error {
	err := prokerDB.db.Create(&proker).Error
	if err != nil {
		return err
	}
	return nil
}

func (p *prokerDatabase) GetByID(id uint) (*domain.EntitasProker, error) {
	var proker domain.EntitasProker
	err := prokerDB.db.First(&proker, id).Error
	if err != nil {
		return nil, err
	}
	return &proker, nil
}

func (p *prokerDatabase) Delete(id uint) error {
	err := prokerDB.db.Where("id = ?", id).Delete(&domain.EntitasProker{}).Error
	if err != nil {
		return err
	}
	return nil
}
