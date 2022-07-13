package route_delivery

import (
	databaseService "BE_WEB_BEM_Proker/implementation/proker/db"
	handlerProker "BE_WEB_BEM_Proker/implementation/proker/handler"
	"BE_WEB_BEM_Proker/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitProkerRouter(g *gin.RouterGroup, db *gorm.DB) {
	dbServ := databaseService.InitProkerDB(db)
	hProker := handlerProker.NewHandlerProker(dbServ)
	g.GET("/", hProker.GetAll)
	g.GET("/:id", hProker.GetByID)
	g.DELETE("/:id", middleware.ValidateJWToken(), hProker.Delete)
	g.POST("/add", middleware.ValidateJWToken(), hProker.Create)
}
