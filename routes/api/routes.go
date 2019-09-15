package api

import (
	"github.com/gin-gonic/gin"
	"github.com/somprabhsharma/the-lazy-traveler/handlers/flightpath"
	"github.com/somprabhsharma/the-lazy-traveler/models"
)

const (
	// BaseURL .
	BaseURL = "/the-lazy-traveler/api/1.0"
)

//Register function registers the APIs to router
func Register(router *gin.Engine, dao *models.Dao) {
	// initialize flight path handler
	flightPathHandler := flightpath.NewHandler(dao)

	lazyJackRoutes := router.Group(BaseURL + "/lazy_jack")
	lazyJackRoutes.POST("", flightPathHandler.ValidateLazyJackRequest, flightPathHandler.FindShortestFlightPath)
}
