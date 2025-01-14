package main

import (
	"database/sql"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/go-kipi/Web-Scraper-Application-Objective/backend/db"
)

var GlobalClientDb *sql.DB

func main() {
	var err error
	dbInterface := db.NewDb()

	GlobalClientDb, err = dbInterface.InitDb()
	if err != nil {
		panic(err)
	}

	router := gin.Default()
	//router.Use(nocache())
	router.Use(cors.New(cors.Config{
		AllowOrigins:  []string{"*"},
		AllowMethods:  []string{"GET", "POST"},
		AllowHeaders:  []string{"Origin", "text/event-stream"},
		ExposeHeaders: []string{"Content-Length"},

		//AllowCredentials: true,
		//AllowOriginFunc: func(origin string) bool {
		//	return origin == "https://github.com"
		//},
		//MaxAge: 12 * time.Hour,
	}))
	router.POST("/scrape", scrapeHandler)
	router.GET("/pokemons", test)
	if err := router.Run(":7777"); err != nil {
		panic(err)
	}
}
