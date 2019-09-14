package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/heroku/x/hmetrics/onload"
	"github.com/somprabhsharma/the-lazy-traveler/constants"
	"log"
	"net/http"
)

func main() {
	router := gin.New()
	router.Use(gin.Logger())

	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello World!")
	})

	err := router.Run(":" + constants.Env.Port)
	if err != nil {
		log.Fatal("Unable to start server")
	}
}
