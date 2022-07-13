package route_delivery

import (
	servAdmin "BE_WEB_BEM_Proker/implementation/admin/db"
	handlerAdmin "BE_WEB_BEM_Proker/implementation/admin/handler"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitAdminRouter(g *gin.RouterGroup, db *gorm.DB) {
	dbServ := servAdmin.InitAdminDB(db)
	hAdmin := handlerAdmin.InitHandlerAdmin(dbServ)
	g.POST("/register", hAdmin.Register)
	g.POST("/login", hAdmin.Login)
	g.GET("/", hAdmin.GetAllAdmin)
	g.GET("/:id", hAdmin.GetAdminByID)
	g.GET("/proker/:id", hAdmin.GetAdminByIDWithProker)
}
