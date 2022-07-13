package db

import (
	"BE_WEB_BEM_Proker/domain"

	"gorm.io/gorm"
)

type adminDatabase struct {
	db *gorm.DB
}

var adminDB *adminDatabase

type DatabaseadminService interface {
	CreateAdmin(admin domain.EntitasAdmin) error
	Login(username string) (*domain.EntitasAdmin, error)
	GetAllAdmin() ([]domain.EntitasAdmin, error)
	GetAdminByIDWithProker(id uint) (*domain.EntitasAdmin, error)
	GetAdminByID(id uint) (*domain.EntitasAdmin, error)
}

func InitAdminDB(db *gorm.DB) DatabaseadminService {
	if adminDB == nil {
		adminDB = &adminDatabase{db}
	}
	return adminDB
}

func (a *adminDatabase) CreateAdmin(admin domain.EntitasAdmin) error {
	err := adminDB.db.Create(&admin).Error
	if err != nil {
		return err
	}
	return nil
}

func (a *adminDatabase) Login(username string) (*domain.EntitasAdmin, error) {
	var admin domain.EntitasAdmin
	err := adminDB.db.Where("username = ?", username).First(&admin).Error
	if err != nil {
		return nil, err
	}
	return &admin, nil
}

func (a *adminDatabase) GetAllAdmin() ([]domain.EntitasAdmin, error) {
	var admins []domain.EntitasAdmin
	err := adminDB.db.Preload("EntitasProkers").Find(&admins).Error
	if err != nil {
		return nil, err
	}
	return admins, nil
}

func (a *adminDatabase) GetAdminByIDWithProker(id uint) (*domain.EntitasAdmin, error) {
	var admin domain.EntitasAdmin
	err := adminDB.db.Preload("EntitasProkers").Where("id = ?", id).First(&admin).Error
	if err != nil {
		return nil, err
	}
	return &admin, nil
}

func (a *adminDatabase) GetAdminByID(id uint) (*domain.EntitasAdmin, error) {
	var admin domain.EntitasAdmin
	err := adminDB.db.Where("id = ?", id).First(&admin).Error
	if err != nil {
		return nil, err
	}
	return &admin, nil
}
