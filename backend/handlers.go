package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-kipi/Web-Scraper-Application-Objective/backend/scrape"
	"net/http"
)

func scrapeHandler(c *gin.Context) {
	//input from req
	type Request struct {
		URL string `json:"url"`
	}
	var request Request
	if err := c.ShouldBindBodyWithJSON(&request); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error":  err.Error(),
			"status": http.StatusInternalServerError,
		})
	}

	//url := "https://pokemondb.net/pokedex/all"
	if result, err := scrape.Run(request.URL); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error":  err.Error(),
			"status": http.StatusInternalServerError,
		})
		return
	} else {
		if errGlobalClientDb := GlobalClientDb.SaveToSqlite(result); errGlobalClientDb != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error":  errGlobalClientDb.Error(),
				"status": http.StatusInternalServerError,
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"status": http.StatusOK,
		})
	}

}

func pokemons(c *gin.Context) {

	count, err := GlobalClientDb.CountFromSqlite()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error":  err.Error(),
			"status": http.StatusInternalServerError,
		})
		return
	}
	p, err := GlobalClientDb.GetFromSqlite()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error":  err.Error(),
			"status": http.StatusInternalServerError,
		})
		return
	}

	var res struct {
		Count    int64       `json:"count"`
		Pokemons interface{} `json:"pokemons"`
	}
	res.Count = count
	res.Pokemons = p
	c.JSON(http.StatusOK, gin.H{
		"data":   res,
		"status": http.StatusOK,
	})
}
