package main

import (
	"fmt"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/realbucksavage/miniauth/api"
	"github.com/realbucksavage/miniauth/database"
)

func main() {

	db, _ := database.Initialize()

	r := gin.Default()
	r.Use(static.Serve("/", static.LocalFile("./static", true)))
	database.Inject(db)

	api.ApplyRoutes(r)

	if err := r.Run(":8080"); err != nil {
		fmt.Printf("Cannot run Gin server: %v\n", err)
	}
}
