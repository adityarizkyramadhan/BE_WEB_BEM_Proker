package domain

import (
	"gorm.io/gorm"
)

type EntitasProker struct {
	gorm.Model
	NamaProker       string `json:"nama_proker"`
	WaktuTerlaksana  string `json:"waktu_terlaksana"`
	Deskripsi        string `json:"deskripsi"`
	PenanggungJawab  string `json:"penanggung_jawab"`
	KontakPJ         string `json:"kontak_pj"`
	LinkCover        string `json:"link_cover"`
	LinkKegiatan     string `json:"link_kegiatan"`
	LinkKegiatanDua  string `json:"link_kegiatan_dua"`
	LinkKegiatanTiga string `json:"link_kegiatan_tiga"`
	EntitasAdminID   uint   `json:"admin_id"`
}

type EntitasAdmin struct {
	gorm.Model
	Username       string          `gorm:"uniqueIndex" json:"username"`
	Password       string          `json:"-"`
	Kementrian     string          `json:"kementrian"`
	EntitasProkers []EntitasProker `gorm:"foreignkey:EntitasAdminID;references:id" json:"prokers"`
}
