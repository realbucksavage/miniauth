package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/realbucksavage/miniauth/api/auth/realms"
)

func ApplyRoutes(router *gin.RouterGroup) {
	realmsGroup := router.Group("/realms")
	{
		realms.ApplyRoutes(realmsGroup)
	}
}
