package handler

import (
	"BE_WEB_BEM_Proker/domain"
	"BE_WEB_BEM_Proker/middleware"
	"BE_WEB_BEM_Proker/utils/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

type handlerProker struct {
	UseCase domain.ProkerService
}

func NewHandlerProker(useCase domain.ProkerService) domain.ProkerHandler {
	return handlerProker{
		UseCase: useCase,
	}
}

func (h handlerProker) GetAll(c *gin.Context) {
	datas, err := h.UseCase.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ResponseWhenFail("Fail to get proker", err.Error()))
	}
	c.JSON(http.StatusOK, response.ResponseWhenSuccess("Success get proker", datas))
}

func (h handlerProker) GetByID(c *gin.Context) {
	var idpkr idProker
	if err := c.BindUri(&idpkr); err != nil {
		c.JSON(http.StatusUnprocessableEntity, response.ResponseWhenFail("Errorr when bind URI", err.Error()))
		return
	}
	data, err := h.UseCase.GetByID(idpkr.Id)
	if err != nil {
		c.JSON(http.StatusUnauthorized, response.ResponseWhenFail("Fail to connect database", err.Error()))
		return
	}
	c.JSON(http.StatusOK, response.ResponseWhenSuccess("Success", data))
}

func (h handlerProker) Create(c *gin.Context) {
	var input domain.EntitasProkerInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusUnprocessableEntity, response.ResponseWhenFail("Error when bind JSON", err.Error()))
		return
	}
	if err := h.UseCase.Create(&domain.EntitasProker{
		NamaProker:      input.NamaProker,
		WaktuTerlaksana: input.WaktuTerlaksana,
		Deskripsi:       input.Deskripsi,
		PenanggungJawab: input.PenanggungJawab,
		Kementrian:      input.Kementrian,
		KontakPJ:        input.KontakPJ,
		LinkImages:      input.LinkImages,
		LinkImages2:     input.LinkImages2,
	}); err != nil {
		c.JSON(http.StatusInternalServerError, response.ResponseWhenFail("Error when create data in database", err.Error()))
		return
	}
	c.JSON(http.StatusCreated, response.ResponseWhenSuccess("Success when add data to database", nil))
}

type idProker struct {
	Id uint `uri:"id"`
}

func (h handlerProker) Delete(c *gin.Context) {
	var idPrkr idProker
	if err := c.ShouldBindUri(&idPrkr); err != nil {
		c.JSON(http.StatusUnprocessableEntity, response.ResponseWhenFail("Fail to bind uri", err.Error()))
		return
	}
	if err := h.UseCase.Delete(idPrkr.Id); err != nil {
		c.JSON(http.StatusInternalServerError, response.ResponseWhenFail("Fail to delete data", err.Error()))
		return
	}
	c.JSON(http.StatusOK, response.ResponseWhenSuccess("Success to delete data", nil))
}

func (h handlerProker) Login(c *gin.Context) {
	var input domain.AdminInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusUnprocessableEntity, response.ResponseWhenFail("Error when bind JSON", err.Error()))
		return
	}
	data, err := h.UseCase.Login(input.Username, input.Password)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, response.ResponseWhenFail("Fail to login", err.Error()))
		return
	}
	token, err := middleware.GenerateJWToken(data.ID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, response.ResponseWhenFail("Fail to generate token", err.Error()))
		return
	}
	c.JSON(http.StatusOK, response.ResponseWhenSuccess("Success to login", gin.H{
		"token": token,
	}))
}

func (h handlerProker) Register(c *gin.Context) {
	var input domain.AdminInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusUnprocessableEntity, response.ResponseWhenFail("Error when bind JSON", err.Error()))
		return
	}
	data, err := h.UseCase.Register(&domain.Admin{
		Username: input.Username,
		Password: input.Password,
	})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, response.ResponseWhenFail("Fail to register", err.Error()))
		return
	}
	token, err := middleware.GenerateJWToken(data.ID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, response.ResponseWhenFail("Fail to generate token", err.Error()))
		return
	}
	c.JSON(http.StatusCreated, response.ResponseWhenSuccess("Success to register", gin.H{
		"token": token,
	}))

}
