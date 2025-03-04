package controllers

import (
	"errors"
	"net/http"

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

// store a post
func StorePost(c *gin.Context) {
	// validate input
	var input ValidationPostInput
	if err := c.ShouldBindJSON(&input); err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			out := make([]ErrorMsg, len(ve))
			for i, fe := range ve {
				out[i] = ErrorMsg{fe.Field(), GetErrorMsg(fe)}
			}
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": out})
		}

		return
	}

	// create post
	post := models.Post{
		Title:   input.Title,
		Content: input.Content,
	}
	models.DB.Create(&post)

	// return json
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Post Created Successfully",
		"data":    post,
	})
}

// get posts by id
func FindPostsById(c *gin.Context) {
	var post models.Post
	if err := models.DB.Where("id = ?", c.Param("id")).First(&post).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
		return
	}

	c.JSON(200, gin.H{
		"success": true,
		"message": "Detail Data Post By Id : " + c.Param("id"),
		"data":    post,
	})
}
