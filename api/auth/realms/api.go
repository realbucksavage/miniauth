package realms

import "github.com/gin-gonic/gin"

func ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func ApplyRoutes(router *gin.RouterGroup) {
	router.GET("/ping", ping)
}
