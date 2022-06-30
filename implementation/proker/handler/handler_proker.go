package handler

import (
	"BE_WEB_BEM_Proker/domain"
	"BE_WEB_BEM_Proker/middleware"
	"BE_WEB_BEM_Proker/utils/response"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"BE_WEB_BEM_Proker/utils"

	"github.com/gin-gonic/gin"
	storage_go "github.com/supabase-community/storage-go"
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
	if data == nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, response.ResponseWhenFail("Username or password is wrong", nil))
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

func (h handlerProker) UploadImage(c *gin.Context) {
	idFoto := c.Param("id")
	// convert id to int
	id, err := strconv.Atoi(idFoto)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, response.ResponseWhenFail("Error when convert id to int", err.Error()))
		return
	}
	fileInput, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, response.ResponseWhenFail("Error when get file", err.Error()))
		return
	}
	file, err := fileInput.Open()
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, response.ResponseWhenFail("Error when open file", err.Error()))
		return
	}
	client := storage_go.NewClient("https://jgjyjvyldoamqndazixl.supabase.co/storage/v1", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJzdXBhYmFzZSIsInJlZiI6ImpnanlqdnlsZG9hbXFuZGF6aXhsIiwicm9sZSI6ImFub24iLCJpYXQiOjE2NDc4MzQ0MDQsImV4cCI6MTk2MzQxMDQwNH0.WVMjJIRoK_cnyfRXdYvTokNWBCCqLWfbeu7xXeZrs6I", nil)
	fileName := fmt.Sprintf("data%d%s", id, fileInput.Filename)
	fileName = strings.ReplaceAll(fileName, ".", "")
	fmt.Println("Sampai selesai file name")
	resp := client.UploadFile("foto-proker", fileName, file)
	fmt.Println(resp)
	fmt.Println("Sampai selesai upload file")
	linkImage := utils.GenerateLinkImage(fileName)
	if err := h.UseCase.SaveImage(linkImage, uint(id)); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, response.ResponseWhenFail("Error when save image", err.Error()))
		return
	}
	c.JSON(http.StatusCreated, response.ResponseWhenSuccess("Success to upload image", nil))
}
