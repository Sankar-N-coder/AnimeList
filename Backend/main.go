package main

import (
	controller "basic/Controller"
	initializers "basic/Initializers"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.Loadenv()
	initializers.Connectdatabase()
}
func main() {
	app := gin.Default()
	app.Use(corsMiddleware())
	app.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello Everyone",
		})
	})
	app.POST("/signup",controller.Register)
	app.Run()
}
func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
