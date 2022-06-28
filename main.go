package main

import (
	"BE_WEB_BEM_Proker/infrastructure/database_connection"
	"BE_WEB_BEM_Proker/infrastructure/database_driver"
	"BE_WEB_BEM_Proker/route"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	e := gin.Default()
	e.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))
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
