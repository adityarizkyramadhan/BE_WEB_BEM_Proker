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
	e.Use(middleware.Cors())
	env, err := database_driver.ReadEnvSupabase()
	if err != nil {
		panic(err)
	}
	db, err := database_connection.MakeConnection(env)
	if err != nil {
		panic(err)
	}
	e.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World!",
		})
	})
	route.RouteAllHandler(e, db)
	if err := e.Run(); err != nil {
		panic(err)
	}
}
