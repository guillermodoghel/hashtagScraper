package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/guillermodoghel/hashtagScraper/models"
	"github.com/guillermodoghel/hashtagScraper/services"
)

//GetContent process the request and call the corresponding services
func GetContent(c *gin.Context) {
	hashtag := c.Param("hashtag")

	instagramFuture := make(chan []models.Post)
	twitterFuture := make(chan []models.Post)

	go func() {
		instagramFuture <- services.GetInstagram(hashtag)
	}()
	go func() {
		twitterFuture <- services.GetTwitter(hashtag)
	}()

	c.JSON(200, gin.H{
		"result": []models.NetworkResponse{
			{Name: "instagram", Posts: <-instagramFuture},
			{Name: "twitter", Posts: <-twitterFuture},
		},
	})
}
