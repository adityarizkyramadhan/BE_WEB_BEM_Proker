package admin

import (
	"BE_WEB_BEM_Proker/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

type responseBph struct {
	Kementrian          string `json:"kementrian"`
	Kontak              string `json:"kontak"`
	NamaBPH             string `json:"nama_bph"`
	DeskripsiKementrian string `json:"deskripsi_kementrian"`
	Username            string `json:"username"`
	Id                  uint   `json:"id"`
}

func ReadProfile(c *gin.Context) {
	idUser := c.MustGet("id").(uint)
	dataDb, err := serviceAdmin.FindByID(idUser)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, helper.Response(false, http.StatusInternalServerError, "Fail to querry database", err))
		return
	}
	c.JSON(http.StatusOK, helper.Response(true, http.StatusOK, "Data is Found", &responseBph{
		Kementrian:          dataDb.Kementrian,
		Kontak:              dataDb.Kontak,
		NamaBPH:             dataDb.NamaBPH,
		DeskripsiKementrian: dataDb.DeskripsiKementrian,
		Username:            dataDb.Username,
		Id:                  dataDb.ID,
	}))
}
