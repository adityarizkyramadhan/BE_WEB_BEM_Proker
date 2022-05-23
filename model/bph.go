package model

import "gorm.io/gorm"

type EntitasBPH struct {
	gorm.Model
	Kementrian          string
	Kontak              string
	Password            string
	NamaBPH             string
	DeskripsiKementrian string
	Username            string
	EntitasProkers      []EntitasProker
}
