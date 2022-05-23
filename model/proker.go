package model

import "gorm.io/gorm"

type EntitasProker struct {
	gorm.Model
	EntitasBPHId    uint
	NamaProker      string
	WaktuTerlaksana string
	Deskripsi       string
	PenanggungJawab string
	KontakPJ        string
	LinkPDFProker   string
	LinkDokumentasi string
}
