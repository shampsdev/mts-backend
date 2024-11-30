package server

import (
	"net/http"

	"api.mts.shamps.dev/internal/search"
	"github.com/gin-gonic/gin"
)

// GetAll godoc
// @Summary Retrieve all persons
// @Description Get a list of all persons
// @Tags persons
// @Produce json
// @Success 200 {array} domain.Person
// @Router /persons [get]
func GetAll(engine search.Engine) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, engine.AllPersons())
	}
}
