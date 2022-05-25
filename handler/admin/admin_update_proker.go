package admin

import (
	"BE_WEB_BEM_Proker/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UpdateProker(c *gin.Context) {
	// idAdmin := c.MustGet("id").(uint)
	var input addProker
	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, helper.Response(false, http.StatusUnprocessableEntity, "Entinity not completed", err))
		return
	}
	// data, err := seviceProker.Update(&model.EntitasProker{
	// 	EntitasBPHId:    idAdmin,
	// 	NamaProker:      input.NamaProker,
	// 	WaktuTerlaksana: input.WaktuTerlaksana,
	// 	Deskripsi:       input.Deskripsi,
	// 	PenanggungJawab: input.PenanggungJawab,
	// 	KontakPJ:        input.KontakPJ,
	// 	LinkPDFProker:   input.LinkPDFProker,
	// 	LinkDokumentasi: input.LinkDokumentasi,
	// })
	// if err != nil {
	// 	c.AbortWithStatusJSON(http.StatusInternalServerError, helper.Response(false, http.StatusInternalServerError, "Fail to update proker", err))
	// 	return
	// }
	c.JSON(http.StatusOK, helper.Response(true, http.StatusOK, "Success update proker", nil))
}
