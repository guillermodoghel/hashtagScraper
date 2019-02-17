package controllers

import (
	"github.com/gin-gonic/gin"
	. "github.com/guillermodoghel/hashtagScraper/models"
	"github.com/guillermodoghel/hashtagScraper/services"
)

func GetContent(c *gin.Context) {
	hashtag := c.Param("hashtag")

	instagramFuture := make(chan []Post)
	twitterFuture := make(chan []Post)
	go func() {
		instagramFuture <- services.GetInstagram(hashtag)
	}()
	go func() {
		twitterFuture <- services.GetTwitter(hashtag)
	}()
	c.JSON(200, gin.H{
		"result": []NetworkResponse{
			{Name: "instagram", Posts: <-instagramFuture},
			{Name: "twitter", Posts: <-twitterFuture},
		},
	})
}
