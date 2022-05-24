package admin

import (
	"BE_WEB_BEM_Proker/helper"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func UpdateAdmin(c *gin.Context) {
	iduser := c.MustGet("id").(uint)
	dataDb, err := serviceAdmin.FindByID(iduser)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, helper.Response(false, http.StatusInternalServerError, "Fail to querry database", err))
		return
	}
	var input newBph
	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, helper.Response(false, http.StatusUnprocessableEntity, "Entinity not completed", err))
		return
	}
	cost, _ := strconv.Atoi(env["BYCRIPT_COST"])
	pass, err := bcrypt.GenerateFromPassword([]byte(input.Password), cost)
	input.Password = string(pass)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, helper.Response(false, http.StatusInternalServerError, "Fail to encrypt password", err))
		return
	}
	dataDb.Kementrian = input.Kementrian
	dataDb.Kontak = input.Kontak
	dataDb.NamaBPH = input.NamaBPH
	dataDb.DeskripsiKementrian = input.Kementrian
	dataDb.Password = input.Password
	dataDb.Username = input.Username
	_, err = serviceAdmin.Update(dataDb)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, helper.Response(false, http.StatusInternalServerError, "Fail to add new admin proker", err))
		return
	}
	c.JSON(http.StatusCreated, helper.Response(true, http.StatusCreated, "Successfully added admin", nil))
}
