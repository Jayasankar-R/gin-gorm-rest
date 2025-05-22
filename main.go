package main

import (
	"github.com/Jayasankar-R/gin-gorm-rest/config"
	"github.com/Jayasankar-R/gin-gorm-rest/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()
	config.Connect()
	// router.GET("/", func(c *gin.Context) {
	// 	c.String(200, "Hello World")
	// })

	routes.UserRoute(router)
	router.Run(":8080")
}
