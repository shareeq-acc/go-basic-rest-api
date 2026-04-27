package main

import (
	"api/db"
	"api/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()

	server.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	routes.RegisterRoutes(server)

	server.Run(":8080") // Listen on port 8080
}
