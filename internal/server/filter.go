package server

import (
	"net/http"

	"api.mts.shamps.dev/internal/search"
	"github.com/gin-gonic/gin"
)

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
