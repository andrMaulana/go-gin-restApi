package controllers

import (
	"github.com/andrMaulana/go-gin-restApi/models"
	"github.com/gin-gonic/gin"
)

func FindPosts(c *gin.Context) {
	// get data from database using models
	var posts []models.Post
	models.DB.Find(&posts)

	// return `json`
	c.JSON(200, gin.H{
		"success": true,
		"message": "List Data Posts",
		"data":    posts,
	})
}
