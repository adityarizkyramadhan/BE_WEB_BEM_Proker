package handler

import (
	"BE_WEB_BEM_Proker/helper"
	"BE_WEB_BEM_Proker/middleware"
	"BE_WEB_BEM_Proker/model"
	"BE_WEB_BEM_Proker/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
)

type newBph struct {
	Kementrian          string `json:"kementrian" binding:"required"`
	Kontak              string `json:"kontak" binding:"required"`
	Password            string `json:"password" binding:"required"`
	NamaBPH             string `json:"nama_bph" binding:"required"`
	DeskripsiKementrian string `json:"deskripsi_kementrian" binding:"required"`
	Username            string `json:"username" binding:"required"`
}

type responseBphLogin struct {
	Token string `json:"token"`
}

var env, _ = godotenv.Read()
var serviceAdmin = service.NewAdminService()

func RegisterAdmin(c *gin.Context) {
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
	data, err := serviceAdmin.Create(&model.EntitasBPH{
		Username:            input.Username,
		Kementrian:          input.Kementrian,
		Kontak:              input.Kontak,
		Password:            input.Password,
		NamaBPH:             input.NamaBPH,
		DeskripsiKementrian: input.DeskripsiKementrian,
	})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, helper.Response(false, http.StatusInternalServerError, "Fail to add new admin proker", err))
		return
	}
	token, err := middleware.GenerateJWToken(data.ID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, helper.Response(false, http.StatusInternalServerError, "Fail to generate jwt token", err))
		return
	}
	c.JSON(http.StatusCreated, helper.Response(true, http.StatusCreated, "Successfully added admin", responseBphLogin{
		Token: token,
	}))
}

type adminLogin struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func LoginAdmin(c *gin.Context) {
	var input adminLogin
	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, helper.Response(false, http.StatusUnprocessableEntity, "Entinity not completed", err))
		return
	}
	data, err := serviceAdmin.FindByUsername(input.Username)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, helper.Response(false, http.StatusInternalServerError, "Error when querry database", err))
		return
	}
	if err := bcrypt.CompareHashAndPassword([]byte(data.Password), []byte(input.Password)); err != nil {
		c.AbortWithStatusJSON(http.StatusForbidden, helper.Response(false, http.StatusUnprocessableEntity, "Password not match", err))
		return
	}
	token, err := middleware.GenerateJWToken(data.ID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, helper.Response(false, http.StatusInternalServerError, "Fail to generate jwt token", err))
		return
	}
	c.JSON(http.StatusOK, helper.Response(true, http.StatusOK, "Successfully login as admin", responseBphLogin{
		Token: token,
	}))
}

type responseBph struct {
	Kementrian          string `json:"kementrian"`
	Kontak              string `json:"kontak"`
	NamaBPH             string `json:"nama_bph"`
	DeskripsiKementrian string `json:"deskripsi_kementrian"`
	Username            string `json:"username"`
}
