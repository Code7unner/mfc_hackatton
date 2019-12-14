package main

import (
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/mfc_hackatton/db"
	"github.com/mfc_hackatton/parser"
	"github.com/mfc_hackatton/scheduler"
	"log"
	"net/http"
	"os"
)

func main() {
	database := db.Connect()
	storage := db.NewStorage(database)

	r := gin.New()

	// Middleware
	r.Use(gin.Logger())

	// Frontend
	r.Use(static.Serve("/", static.LocalFile("client/build", true)))
	r.LoadHTMLGlob("client/build/*.html")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	// Parsing .xlsx file
	r.GET("/api/parser", gin.WrapF(parser.Parse))
	// Download and update MFC information in db
	r.GET("/api/server", gin.WrapF(storage.GetServerStats))
	// Download and update MFC statistics in db
	r.GET("/api/statistics", gin.WrapF(storage.GetStatistics))

	// Port
	if err := r.Run(":" + os.Getenv("PORT")); err != nil {
		log.Fatal(err)
	}

	// Run the scheduler
	scheduler.Schedule(storage)
}
