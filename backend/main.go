package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	SetupRoutes(r)
	r.Run(":8080")
}

func SetupRoutes(r *gin.Engine) {
	api := r.Group("/api")
	{
		api.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{"message": "pong"})
		})
		
	}
}