package server

import (
	"net/http"

	"api.mts.shamps.dev/internal/search"
	"github.com/gin-gonic/gin"
)

// GetAll godoc
// @Summary Retrieve all divisions
// @Description Get a list of all divisions
// @Tags persons
// @Produce json
// @Success 200 {array} string
// @Router /persons/divisions [get]
func GetAllDivisions(engine search.Engine) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, engine.AllDivisions())
	}
}
