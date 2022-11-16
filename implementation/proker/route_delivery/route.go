package route_delivery

import (
	databaseService "BE_WEB_BEM_Proker/implementation/proker/db"
	handlerProker "BE_WEB_BEM_Proker/implementation/proker/handler"
	"BE_WEB_BEM_Proker/middleware"
	"time"

	"github.com/gin-contrib/cache"
	"github.com/gin-contrib/cache/persistence"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitProkerRouter(g *gin.RouterGroup, db *gorm.DB) {
	store := persistence.NewInMemoryStore(10 * time.Minute)
	dbServ := databaseService.InitProkerDB(db)
	hProker := handlerProker.NewHandlerProker(dbServ)
	g.GET("/", cache.CachePage(store, 10*time.Minute, hProker.GetAll))
	g.GET("/:id", cache.CachePage(store, 10*time.Minute, hProker.GetByID))
	g.DELETE("/:id", middleware.ValidateJWToken(), hProker.Delete)
	g.POST("/add", middleware.ValidateJWToken(), hProker.Create)
	g.GET("/paging", cache.CachePage(store, 10*time.Minute, hProker.Paging))
}
