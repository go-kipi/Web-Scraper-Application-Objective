package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/go-kipi/Web-Scraper-Application-Objective/backend/db"
	"os"
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
	router.Use(cors.New(cors.Config{
		AllowOrigins:  []string{"*"},
		AllowMethods:  []string{"GET", "POST"},
		ExposeHeaders: []string{"Content-Length"},
	}))
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
