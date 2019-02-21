package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/guillermodoghel/hashtagScraper/controllers"
)

func main() {
	fmt.Println("Hello, 世界")

	r := gin.Default()
	r.Static("/www", "./www")
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/search/:hashtag", controllers.GetContent)
	r.GET("/stream/:hashtag", controllers.StreamContent)

	r.Run()
}
