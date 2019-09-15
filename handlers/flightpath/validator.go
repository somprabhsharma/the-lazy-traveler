package flightpath

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/somprabhsharma/the-lazy-traveler/constants/errorconsts"
	"github.com/somprabhsharma/the-lazy-traveler/constants/literals"
	"github.com/somprabhsharma/the-lazy-traveler/entities/flightpath"
	"github.com/somprabhsharma/the-lazy-traveler/utils/logger"
	"net/http"
)

// ValidateLazyJackRequest validate request body in lazy jack apis by trying to bind it
func (h *Handler) ValidateLazyJackRequest(c *gin.Context) {
	var lazyJackRequest flightpath.LazyJackRequest

	if err := c.Bind(&lazyJackRequest); err != nil {
		logger.Err(literals.LazyJack, "error in binding request", err, lazyJackRequest)
		_ = c.AbortWithError(http.StatusBadRequest, errors.New(errorconsts.InvalidRequest))
	}

	c.Set("lazyJackRequest", lazyJackRequest)
}
