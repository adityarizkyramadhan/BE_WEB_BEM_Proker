package main

import (
	"BE_WEB_BEM_Proker/infrastructure/database_connection"
	"BE_WEB_BEM_Proker/infrastructure/database_driver"
	"BE_WEB_BEM_Proker/middleware"
	"BE_WEB_BEM_Proker/route"

	"github.com/gin-gonic/gin"
)

func main() {
	e := gin.Default()
	env, err := database_driver.ReadEnvSupabase()
	if err != nil {
		panic(err)
	}
	db, err := database_connection.MakeConnection(env)
	if err != nil {
		panic(err)
	}
	e.Use(middleware.CORS())
	e.GET("", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World!",
			"test":    "Ini lagi mencoba cd ke empat menambahi cache hahaha",
		})
	})
	route.InitRouteAll(e, db)
	if err := e.Run(":8070"); err != nil {
		panic(err)
	}
}
