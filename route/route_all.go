package route

import (
	"BE_WEB_BEM_Proker/implementation/proker/handler"
	"BE_WEB_BEM_Proker/implementation/proker/route_delivery"
	"BE_WEB_BEM_Proker/implementation/proker/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RouteAllHandler(e *gin.Engine, db *gorm.DB) {
	serv := service.NewServiceProker(db)
	handler := handler.NewHandlerProker(serv)
	route_delivery.NewRouter(e, handler)
}
