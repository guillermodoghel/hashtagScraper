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
	facebookFuture := make(chan []Post)

	go func() {
		instagramFuture <- services.GetInstagram(hashtag)
	}()
	go func() {
		twitterFuture <- services.GetTwitter(hashtag)
	}()
	go func() {
		facebookFuture <- services.GetFacebook(hashtag)
	}()
	c.JSON(200, gin.H{
		"result": []NetworkResponse{
			{Name: "facebook", Posts: <-facebookFuture},
			{Name: "instagram", Posts: <-instagramFuture},
			{Name: "twitter", Posts: <-twitterFuture},
		},
	})
}
