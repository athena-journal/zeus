package main

import (
	"github.com/gin-gonic/gin"

	mongodb "github.com/athena-journal/zeus/mongodb"
)

func main() {
	r := gin.Default()

	r.GET("/users", func(c *gin.Context) {
		documents := mongodb.FindAll("users") // Get All User Documents

		c.JSON(200, gin.H{
			"data": documents,
		})
	})

	r.Run(":8080")
}
