package main

import (
	"github.com/andrMaulana/go-gin-restApi/controllers"
	"github.com/andrMaulana/go-gin-restApi/models"
	"github.com/gin-gonic/gin"
)

func main() {
	// inisialisasi gin
	router := gin.Default()

	// panggil koneksi database
	models.ConnectDatabase()
	// membuat route dengan method get
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World!",
		})
	})

	router.GET("/api/posts", controllers.FindPosts)

	router.Run(":3000")
}
