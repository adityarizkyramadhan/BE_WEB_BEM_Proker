package route_delivery

import (
	"BE_WEB_BEM_Proker/domain"
	"BE_WEB_BEM_Proker/middleware"

	"github.com/gin-gonic/gin"
)

func NewRouter(e *gin.Engine, s domain.ProkerHandler) {
	e.GET("/proker", s.GetAll)
	e.GET("/proker/:id", s.GetByID)
	e.POST("/proker", middleware.ValidateJWToken(), s.Create)
	e.DELETE("/proker/:id", middleware.ValidateJWToken(), s.Delete)
	e.POST("/admin/login", s.Login)
}
