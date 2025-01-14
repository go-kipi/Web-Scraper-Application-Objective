package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-kipi/Web-Scraper-Application-Objective/backend/scrape"
	"net/http"
)

func scrapeHandler(c *gin.Context) {
	//input from req
	url := "https://pokemondb.net/pokedex/all"
	if result, err := scrape.Run(url); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
	} else {
		//save result on DB
		c.JSON(http.StatusOK, gin.H{
			"data": result,
		})
	}

}

func test(c *gin.Context) {
	x := GlobalClientDb
	err := x.Ping()
	fmt.Println(err)
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
