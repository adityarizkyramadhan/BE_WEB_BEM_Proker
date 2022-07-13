package handler

import (
	"BE_WEB_BEM_Proker/domain"
	"BE_WEB_BEM_Proker/implementation/admin/db"
	"BE_WEB_BEM_Proker/middleware"
	"BE_WEB_BEM_Proker/utils/response"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

var handler *handlerAdmin

type handlerAdmin struct {
	db db.DatabaseadminService
}

type HandlerAdmin interface {
	Register(c *gin.Context)
	Login(c *gin.Context)
	GetAllAdmin(c *gin.Context)
	GetAdminByID(c *gin.Context)
	GetAdminByIDWithProker(c *gin.Context)
}

func InitHandlerAdmin(db db.DatabaseadminService) HandlerAdmin {
	if handler == nil {
		handler = &handlerAdmin{db}
	}
	return handler
}

type inputAdminRegister struct {
	Username   string `json:"username"`
	Password   string `json:"password"`
	Kementrian string `json:"kementrian"`
}

func (h *handlerAdmin) Register(c *gin.Context) {
	var dataInput inputAdminRegister
	if err := c.BindJSON(&dataInput); err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, response.ResponseWhenFail("Errorr when bind JSON", err.Error()))
		return
	}
	password, err := bcrypt.GenerateFromPassword([]byte(dataInput.Password), bcrypt.DefaultCost)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, response.ResponseWhenFail("Error when generate password", err.Error()))
		return
	}
	data := domain.EntitasAdmin{
		Username:   dataInput.Username,
		Password:   string(password),
		Kementrian: dataInput.Kementrian,
	}
	if err := h.db.CreateAdmin(data); err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, response.ResponseWhenFail("Error when create admin", err.Error()))
		return
	}
	token, err := middleware.GenerateJWToken(data.ID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, response.ResponseWhenFail("Error when generate token", err.Error()))
		return
	}
	c.JSON(http.StatusOK, response.ResponseWhenSuccess("Success create admin", gin.H{
		"token": token,
	}))
}

func (h *handlerAdmin) Login(c *gin.Context) {
	var dataInput inputAdminRegister
	if err := c.BindJSON(&dataInput); err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, response.ResponseWhenFail("Errorr when bind JSON", err.Error()))
		return
	}
	data, err := h.db.Login(dataInput.Username)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, response.ResponseWhenFail("Fail to connect database", err.Error()))
		return
	}
	if err := bcrypt.CompareHashAndPassword([]byte(data.Password), []byte(dataInput.Password)); err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, response.ResponseWhenFail("Fail to connect database", err.Error()))
		return
	}
	token, err := middleware.GenerateJWToken(data.ID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, response.ResponseWhenFail("Error when generate token", err.Error()))
		return
	}
	c.JSON(http.StatusOK, response.ResponseWhenSuccess("Success login", gin.H{
		"token": token,
	}))
}

func (h *handlerAdmin) GetAllAdmin(c *gin.Context) {
	data, err := h.db.GetAllAdmin()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, response.ResponseWhenFail("Error when get all admin", err.Error()))
		return
	}
	c.JSON(http.StatusOK, response.ResponseWhenSuccess("Success get all admin", data))
}

func (h *handlerAdmin) GetAdminByID(c *gin.Context) {
	id := c.Param("id")
	//convert string to uint
	idAdmin, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, response.ResponseWhenFail("Error when convert string to uint", err.Error()))
		return
	}
	data, err := h.db.GetAdminByID(uint(idAdmin))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, response.ResponseWhenFail("Error when get admin by id", err.Error()))
		return
	}
	c.JSON(http.StatusOK, response.ResponseWhenSuccess("Success get admin by id", data))
}

func (h *handlerAdmin) GetAdminByIDWithProker(c *gin.Context) {
	id := c.Param("id")
	//convert string to uint
	idAdmin, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, response.ResponseWhenFail("Error when convert string to uint", err.Error()))
		return
	}
	data, err := h.db.GetAdminByIDWithProker(uint(idAdmin))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, response.ResponseWhenFail("Error when get admin by id", err.Error()))
		return
	}
	c.JSON(http.StatusOK, response.ResponseWhenSuccess("Success get admin by id", data))
}
