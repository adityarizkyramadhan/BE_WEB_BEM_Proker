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
	LinkPDFProker   string
	LinkDokumentasi string
}

type EntitasProkerInput struct {
	NamaProker      string `json:"nama_proker,omitempty" binding:"required"`
	WaktuTerlaksana string `json:"waktu_terlaksana,omitempty" binding:"required"`
	Kementrian      string `json:"kementrian" binding:"required"'`
	Deskripsi       string `json:"deskripsi,omitempty" binding:"required"`
	PenanggungJawab string `json:"penanggung_jawab,omitempty" binding:"required"`
	KontakPJ        string `json:"kontak_pj,omitempty" binding:"required"`
	LinkPDFProker   string `json:"link_pdf_proker,omitempty" binding:"required"`
	LinkDokumentasi string `json:"link_dokumentasi,omitempty" binding:"required"`
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
	Delete(*gin.Context)
}
