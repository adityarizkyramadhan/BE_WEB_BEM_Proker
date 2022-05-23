package service

import "BE_WEB_BEM_Proker/model"

type serviceAdmin interface {
	Create(*model.EntitasBPH) (*model.EntitasBPH, error)
	Read() (*[]model.EntitasBPH, error)
	Update(*model.EntitasBPH) (*model.EntitasBPH, error)
	Delete(*model.EntitasBPH) (*model.EntitasBPH, error)
	FindByUsername(string) (*model.EntitasBPH, error)
	FindByID(uint) (*model.EntitasBPH, error)
}

type adminService struct{}

func NewAdminService() serviceAdmin {
	return &adminService{}
}
func (s *adminService) FindByID(id uint) (*model.EntitasBPH, error) {
	var entitas model.EntitasBPH
	if err := DB.Where("id = ?", id).Take(&entitas).Error; err != nil {
		return nil, err
	}
	return &entitas, nil
}

func (s *adminService) FindByUsername(username string) (*model.EntitasBPH, error) {
	var entitas model.EntitasBPH
	if err := DB.Where("username = ?", username).Take(&entitas).Error; err != nil {
		return nil, err
	}
	return &entitas, nil
}

func (s *adminService) Create(entitas *model.EntitasBPH) (*model.EntitasBPH, error) {
	if err := DB.Create(&entitas).Error; err != nil {
		return nil, err
	}
	return entitas, nil
}

func (s *adminService) Update(entitas *model.EntitasBPH) (*model.EntitasBPH, error) {
	if err := DB.Save(&entitas).Error; err != nil {
		return nil, err
	}
	return entitas, nil
}

func (s *adminService) Delete(entitas *model.EntitasBPH) (*model.EntitasBPH, error) {
	if err := DB.Delete(&entitas).Error; err != nil {
		return nil, err
	}
	return entitas, nil
}

func (s *adminService) Read() (*[]model.EntitasBPH, error) {
	var entitasBanyak *[]model.EntitasBPH
	if err := DB.Find(&entitasBanyak).Error; err != nil {
		return nil, err
	}
	return entitasBanyak, nil
}
