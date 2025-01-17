package main

import "github.com/gin-gonic/gin"

func main() {
	// inisialisasi gin
	router := gin.Default()

	// membuat route dengan method get
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World!",
		})
	})

	router.Run(":3000")
}
