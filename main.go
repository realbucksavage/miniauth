package main

import (
	"fmt"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/realbucksavage/miniauth/database"
	"github.com/realbucksavage/miniauth/realms"
	"net/http"
)

func main() {

	db := database.Initialize()

	r := gin.Default()
	r.Use(static.Serve("/", static.LocalFile("./static/build", true)))
	r.Use(database.Inject(db))

	realms.ApplyRoutes(r)

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"code": "NOT_FOUND",
		})
	})

	if err := r.Run(":8080"); err != nil {
		fmt.Printf("Cannot run Gin server: %v\n", err)
	}
}
