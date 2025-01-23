package controllers

import (
	"github.com/andrMaulana/go-gin-restApi/models"
	"github.com/gin-gonic/gin"
)

// type validation post input
type ValidationPostInput struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
}

// type error message
type ErrorMsg struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

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
