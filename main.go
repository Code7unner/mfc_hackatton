package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mfc_hackatton/parser"
)

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/api/parser", gin.WrapF(parser.Parse))

	r.Run()
}
