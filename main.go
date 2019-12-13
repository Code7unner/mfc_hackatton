package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mfc_hackatton/db"
	"github.com/mfc_hackatton/parser"
	"log"
)

func main() {
	database := db.Connect()
	st := db.NewStorage(database)

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/api/parser", gin.WrapF(parser.Parse))

	r.GET("/api/server", gin.WrapF(st.GetServerStats))

	if err := r.Run(); err != nil {
		log.Fatal(err)
	}
}
