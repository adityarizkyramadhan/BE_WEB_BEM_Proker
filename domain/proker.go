package domain

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type EntitasProker struct {
	*gorm.Model
	NamaProker      string
	Kementrian      string
	WaktuTerlaksana string
	Deskripsi       string
	PenanggungJawab string
	KontakPJ        string
	LinkImages      string
	LinkImages2     string
}

type Admin struct {
	*gorm.Model
	Username string `json:"username"`
	Password string `json:"password"`
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
	LinkPdfProker   string `json:"link_pdf_proker,omitempty" binding:"required"`
	LinkImages      string `json:"save_images,omitempty" binding:"required"`
	LinkImages2     string `json:"save_images2,omitempty" binding:"required"`
}

type ProkerService interface {
	GetAll() ([]EntitasProker, error)
	GetByID(uint) (EntitasProker, error)
	Create(*EntitasProker) error
	Update(uint, *EntitasProker) error
	Delete(uint) error
	Login(string, string) (*Admin, error)
	Register(*Admin) (*Admin, error)
}

type ProkerHandler interface {
	GetAll(*gin.Context)
	GetByID(*gin.Context)
	Create(*gin.Context)
	Delete(*gin.Context)
	Login(*gin.Context)
	Register(*gin.Context)
}
