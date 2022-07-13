package route

import (
	routeAdmin "BE_WEB_BEM_Proker/implementation/admin/route_delivery"
	routeProker "BE_WEB_BEM_Proker/implementation/proker/route_delivery"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitRouteAll(e *gin.Engine, db *gorm.DB) {
	routeAdmin.InitAdminRouter(e.Group("/admin"), db)
	routeProker.InitProkerRouter(e.Group("/proker"), db)
}
