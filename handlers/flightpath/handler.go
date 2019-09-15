package flightpath

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/somprabhsharma/the-lazy-traveler/constants/errorconsts"
	"github.com/somprabhsharma/the-lazy-traveler/constants/literals"
	"github.com/somprabhsharma/the-lazy-traveler/controllers/flightpath"
	entities "github.com/somprabhsharma/the-lazy-traveler/entities/flightpath"
	"github.com/somprabhsharma/the-lazy-traveler/models"
	"github.com/somprabhsharma/the-lazy-traveler/utils/logger"
	"net/http"
)

// Handler is a struct which will act like a handler for flight path related APIs
type Handler struct {
	flightPathController flightpath.Controller
}

// NewHandler is a constructor for Handler struct
func NewHandler(dao *models.Dao) *Handler {
	return &Handler{
		flightPathController: *flightpath.NewController(dao),
	}
}

// FindShortestFlightPath finds shortest flight path
func (h *Handler) FindShortestFlightPath(c *gin.Context) {
	v, ok := c.Get("lazyJackRequest")
	if !ok {
		_ = c.AbortWithError(http.StatusBadRequest, errors.New(errorconsts.InvalidRequest))
		return
	}

	body, _ := v.(entities.LazyJackRequest)
	logger.Info(literals.LazyJack, "Request received to find shortest flight path with data", body)

	shortestPath, err := h.flightPathController.FindShortestFlightPath(body)
	if err != nil {
		logger.Err(literals.LazyJack, "Error while finding shortest flight path", err, body)
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{literals.FlightPlan: shortestPath})
}
