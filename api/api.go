package api

import (
	"github.com/gin-gonic/gin"
	"github.com/realbucksavage/miniauth/api/auth"
)

func ApplyRoutes(router *gin.Engine) {
	authGroup := router.Group("/auth")
	{
		auth.ApplyRoutes(authGroup)
	}
}
