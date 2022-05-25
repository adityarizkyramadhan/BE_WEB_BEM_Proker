package admin

import (
	"BE_WEB_BEM_Proker/helper"
	"BE_WEB_BEM_Proker/model"
	"BE_WEB_BEM_Proker/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type addProker struct {
	// EntitasBPHId    uint
	NamaProker      string `json:"nama_proker" binding:"required"`
	WaktuTerlaksana string `json:"waktu_terlaksana" binding:"required"`
	Deskripsi       string `json:"deskripsi" binding:"required"`
	PenanggungJawab string `json:"penanggung_jawab" binding:"required"`
	KontakPJ        string `json:"kontak_pj" binding:"required"`
	LinkPDFProker   string `json:"link_pdf_proker" binding:"required"`
	LinkDokumentasi string `json:"link_dokumentasi" binding:"required"`
}

var seviceProker = service.NewBPHService()

func AddNewProker(c *gin.Context) {
	idAdmin := c.MustGet("id").(uint)
	var input addProker
	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, helper.Response(false, http.StatusUnprocessableEntity, "Entinity not completed", err))
		return
	}
	data, err := seviceProker.Create(&model.EntitasProker{
		EntitasBPHId:    idAdmin,
		NamaProker:      input.NamaProker,
		WaktuTerlaksana: input.WaktuTerlaksana,
		Deskripsi:       input.Deskripsi,
		PenanggungJawab: input.PenanggungJawab,
		KontakPJ:        input.KontakPJ,
		LinkPDFProker:   input.LinkPDFProker,
		LinkDokumentasi: input.LinkDokumentasi,
	})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, helper.Response(false, http.StatusInternalServerError, "Fail to add new proker", err))
		return
	}
	c.JSON(http.StatusOK, helper.Response(true, http.StatusOK, "Success add new proker", data))
}
