package api

import "github.com/gin-gonic/gin"

func ApplyRoute(r *gin.RouterGroup) {
	r.POST("/", create)

	r.GET("/", list)
	r.GET("/:name", findByName)

	r.DELETE("/:name", remove)
}
