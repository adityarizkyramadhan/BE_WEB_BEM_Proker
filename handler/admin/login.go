package admin

import (
	"BE_WEB_BEM_Proker/helper"
	"BE_WEB_BEM_Proker/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

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
