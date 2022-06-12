package domain

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type EntitasProker struct {
	*gorm.Model
	NamaProker      string
	WaktuTerlaksana string
	Deskripsi       string
	PenanggungJawab string
	KontakPJ        string
	LinkPDFProker   string
	LinkDokumentasi string
}

type ProkerService interface {
	GetAll() ([]EntitasProker, error)
	GetByID(uint) (EntitasProker, error)
	Create(*EntitasProker) error
	Update(uint, *EntitasProker) error
	Delete(uint) error
}

type ProkerHandler interface {
	GetAll(*gin.Context)
	GetByID(*gin.Context)
	Create(*gin.Context)
	Update(*gin.Context)
	Delete(*gin.Context)
}
