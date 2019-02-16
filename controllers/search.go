package controllers

import (
	"github.com/gin-gonic/gin"
	. "github.com/guillermodoghel/example/models"
	"github.com/guillermodoghel/example/services"
)

func GetContent(c *gin.Context) {
	hashtag := c.Param("hashtag")

	c.JSON(200, gin.H{
		"result": []NetworkResponse{
			{Name: "instagram", Posts: services.GetInstagram(hashtag)},
		},
	})
}
