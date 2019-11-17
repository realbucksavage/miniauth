package realms

import (
	"github.com/gin-gonic/gin"
	"github.com/realbucksavage/miniauth/realms/api"
)

func ApplyRoutes(r *gin.Engine) {
	realmsRoute := r.Group("/realms")
	{
		realmsRoute.POST("/init-master-realm", initMasterRealm)
		api.ApplyRoute(realmsRoute)
	}
}
