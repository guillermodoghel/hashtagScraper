package controllers

import (
	"log"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	. "github.com/guillermodoghel/hashtagScraper/models"
	"github.com/guillermodoghel/hashtagScraper/services"
	"github.com/guillermodoghel/hashtagScraper/utils"
)

//GetContent process the request and call the corresponding services
func GetContent(c *gin.Context) {
	start := time.Now()
	hashtag := c.Param("hashtag")

	instagramFuture := make(chan []Post)
	twitterFuture := make(chan []Post)

	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		tempInstagram := services.GetInstagram(hashtag)
		utils.SignPosts(&tempInstagram)
		instagramFuture <- tempInstagram
		wg.Done()
	}()

	go func() {
		tempTwitter := services.GetTwitter(hashtag)
		utils.SignPosts(&tempTwitter)
		twitterFuture <- tempTwitter
		wg.Done()
	}()

	go func() {
		wg.Wait()
		close(instagramFuture)
		close(twitterFuture)
	}()

	c.JSON(200, gin.H{
		"result": []NetworkResponse{
			{Name: "instagram", Posts: <-instagramFuture},
			{Name: "twitter", Posts: <-twitterFuture},
		},
	})
	elapsed := time.Since(start)
	log.Printf("Scrapping took %s", elapsed)

}
