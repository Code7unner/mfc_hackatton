package main

import (
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	_ "github.com/heroku/x/hmetrics/onload"
	"github.com/mfc_hackatton/db"
	"github.com/mfc_hackatton/parser"
	"github.com/mfc_hackatton/scheduler"
	"log"
	"net/http"
)

func main() {
	//port := ":" + os.Getenv("PORT")

	database := db.Connect()
	storage := db.NewStorage(database)

	r := gin.New()

	// Middleware
	r.Use(gin.Logger())
	r.Use(CORSMiddleware())

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
	// Download and update MFC statistics in db
	r.GET("/api/mfc", gin.WrapF(storage.GetMFCStats))

	// Run the scheduler
	go scheduler.Schedule(storage)

	// Port
	if err := r.Run(":8081"); err != nil {
		log.Fatal(err)
	}
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
