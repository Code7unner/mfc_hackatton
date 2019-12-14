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

	r.LoadHTMLGlob("client/build/*.html")
	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", nil)
	})

	// Parsing .xlsx file
 	r.GET("/api/parser", gin.WrapF(parser.Parse))
	// Download and update MFC information in db
	r.GET("/api/server", gin.WrapF(st.GetServerStats))
	// Download and update MFC statistics in db
	r.GET("/api/statistics", gin.WrapF(st.GetStatistics))

	if err := r.Run(); err != nil {
		log.Fatal(err)
	}
}
