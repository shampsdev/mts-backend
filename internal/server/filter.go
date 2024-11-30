package server

import (
	"net/http"

	"api.mts.shamps.dev/internal/search"
	"github.com/gin-gonic/gin"
)

// Filter godoc
// @Summary Filter persons based on query parameters
// @Description Filters persons based on query parameters like id, name, and status.
// @Tags persons
// @Produce json
// @Param id query string false "Filter by ID"
// @Param name query string false "Filter by Name"
// @Param status query string false "Filter by Status"
// @Success 200 {array} domain.PersonNode
// @Router /persons/filter [get]
func Filter(engine search.Engine) gin.HandlerFunc {
	return func(c *gin.Context) {
		var filters []search.Filter

		if id := c.Query("id"); id != "" {
			filters = append(filters, search.Filter{Key: "id", Val: id})
		}
		if name := c.Query("name"); name != "" {
			filters = append(filters, search.Filter{Key: "name", Val: name})
		}
		if status := c.Query("status"); status != "" {
			filters = append(filters, search.Filter{Key: "status", Val: status})
		}

		c.JSON(http.StatusOK, engine.Filter(filters))
	}
}
