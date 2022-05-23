package handler

import (
	"github.com/gin-gonic/gin"
)

func Handler1(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello World",
	})
	c.SetCookie("name", "value", 3600, "/", "localhost", false, true)
}
