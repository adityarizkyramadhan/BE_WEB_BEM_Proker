package domain

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type EntitasProker struct {
	*gorm.Model     `json:"-"`
	NamaProker      string    `json:"nama_proker"`
	Kementrian      string    `json:"kementrian"`
	WaktuTerlaksana string    `json:"waktu_terlaksana"`
	Deskripsi       string    `json:"deskripsi"`
	PenanggungJawab string    `json:"penanggung_jawab"`
	KontakPJ        string    `json:"kontak_pj"`
	LinkImage       LinkImage `gorm:"foreignkey:IdEntitasProker" json:"link_image"`
}

type LinkImage struct {
	gorm.Model      `json:"-"`
	Link            string `json:"link"`
	IdEntitasProker uint   `json:"-"`
}

type Admin struct {
	*gorm.Model
	Username string `gorm:"uniqueIndex"`
	Password string
}

type AdminInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type EntitasProkerInput struct {
	NamaProker      string `json:"nama_proker,omitempty" binding:"required"`
	WaktuTerlaksana string `json:"waktu_terlaksana,omitempty" binding:"required"`
	Kementrian      string `json:"kementrian" binding:"required"`
	Deskripsi       string `json:"deskripsi,omitempty" binding:"required"`
	PenanggungJawab string `json:"penanggung_jawab,omitempty" binding:"required"`
	KontakPJ        string `json:"kontak_pj,omitempty" binding:"required"`
}

type ProkerService interface {
	GetAll() (*[]EntitasProker, error)
	GetByID(uint) (*EntitasProker, error)
	Create(*EntitasProker) error
	Update(uint, *EntitasProker) error
	Delete(uint) error
	Login(string, string) (*Admin, error)
	Register(*Admin) (*Admin, error)
	SaveImage(string, uint) error
}

type ProkerHandler interface {
	GetAll(*gin.Context)
	GetByID(*gin.Context)
	Create(*gin.Context)
	Delete(*gin.Context)
	Login(*gin.Context)
	Register(*gin.Context)
	UploadImage(*gin.Context)
}
