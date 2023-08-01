package main

import (
	"ginapi/middleware"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		c.Set("example", "12345")
		c.Next()
		latency := time.Since(t)
		log.Print(latency)
		status := c.Writer.Status()
		log.Println(status)
	}
}

func main() {
	gin.ForceConsoleColor()
	r := gin.New()
	r.Use(Logger())
	r.Use(middleware.CreateConnection())
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/albums", middleware.GetAlbums)
	r.GET("/albums/:id", middleware.GetAlbumByID)
	r.POST("/albums", middleware.PostAlbums)
	r.Run("localhost:8080")
}
