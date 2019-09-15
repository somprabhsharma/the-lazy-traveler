package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/heroku/x/hmetrics/onload"
	"github.com/somprabhsharma/the-lazy-traveler/constants"
	"github.com/somprabhsharma/the-lazy-traveler/middlewares"
	"github.com/somprabhsharma/the-lazy-traveler/models"
	"github.com/somprabhsharma/the-lazy-traveler/routes/api"
	"log"
	"net/http"
)

func main() {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(middlewares.HandleErrors)

	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello World!")
	})

	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "Pong!")
	})

	// initialize dao layer
	dao := models.NewDao()

	// register api routes
	api.Register(router, dao)

	err := router.Run(":" + constants.Env.Port)
	if err != nil {
		log.Fatal("Unable to start server")
	}
}
