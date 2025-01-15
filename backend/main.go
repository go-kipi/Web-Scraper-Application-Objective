package main

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/go-kipi/Web-Scraper-Application-Objective/backend/db"
	"os"
	"time"
)

var GlobalClientDb db.DbI

func main() {
	var err error
	dbInterface := db.NewDb()

	err = dbInterface.InitDb()
	if err != nil {
		panic(err)
	}
	GlobalClientDb = dbInterface

	router := gin.Default()
	router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))
	//router.TrustedPlatform = gin.PlatformGoogleAppEngine
	//router.TrustedPlatform = "X-CDN-IP"
	router.SetTrustedProxies([]string{})
	//router.RemoteIPHeaders = []string{"X-Forwarded-For", "X-Real-IP"}
	router.Use(cors.Default())
	router.POST("/scrape", scrapeHandler)
	router.GET("/pokemons", pokemons)

	port := os.Getenv("PORT")
	if port == "" {
		port = "7777"
	}
	if err := router.Run(":" + port); err != nil {
		panic(err)
	}
}
