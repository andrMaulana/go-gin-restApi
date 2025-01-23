package controllers

import (
	"github.com/andrMaulana/go-gin-restApi/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
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

// function get error message
func GetErrorMsg(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "This field is required"
	}

	return "Unknow error"
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
