package middlewares

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/somprabhsharma/the-lazy-traveler/constants/errorconsts"
	"github.com/somprabhsharma/the-lazy-traveler/constants/literals"
	"github.com/somprabhsharma/the-lazy-traveler/utils/logger"
	"net/http"
	"strconv"
)

//HandleErrors handles whatever error the API has returned at one place
func HandleErrors(c *gin.Context) {
	// setting Content-Type in response header as application/json
	// we can't set Content-Type afterwards if there is c.Abort.. call in code
	// in such cases Content-Type is returned as text/plain which we don't want
	// therefore we are setting it as application/json at the very initial point
	c.Writer.Header().Set("Content-Type", "application/json; charset=utf-8")

	c.Next() // execute all the handlers

	// at this point, all the handlers finished. Let's read the errors!
	if len(c.Errors) == 0 {
		return
	}
	err := c.Errors.Last()

	// default ltError
	ltError := errorconsts.LTError{
		Message:  errorconsts.GenericErrorMessage,
		Code:     errorconsts.GenericErrorCode,
		HTTPCode: http.StatusBadRequest,
		Err:      err.Error(),
	}

	// if ltErrorMap have the mapping for the received error
	// assign the corresponding ltError to override default
	if ltErr, ok := errorconsts.LTErrorMap[err.Error()]; ok {
		ltError = ltErr

		// assign default values to missing fields
		if ltError.HTTPCode == 0 {
			ltError.HTTPCode = http.StatusBadRequest
		}

		if ltError.Err == "" {
			ltError.Err = err.Error()
		}
	}

	logger.Info(literals.LazyJack, "Code: "+strconv.Itoa(ltError.Code)+" Message: "+ltError.Message+" Error: "+ltError.Err, nil)

	c.AbortWithStatusJSON(ltError.HTTPCode, ltError)
	fmt.Println(c.Writer.Header().Get("Content-Type"))
}
