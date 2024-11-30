package server

import (
	"net/http"

	"api.mts.shamps.dev/internal/search"
	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine) {
	engine := search.NewJSONEngine()

	r.GET("/persons", func(c *gin.Context) {
		c.JSON(http.StatusOK, engine.GetAll())
	})

	r.GET("/persons/filter", func(c *gin.Context) {
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
	})

	r.GET("/persons/search", func(c *gin.Context) {
		text := c.Query("text")
		if text == "" {
			c.JSON(http.StatusBadRequest, gin.H{"message": "search query is empty"})
			return
		}

		c.JSON(http.StatusOK, engine.Search(text))
	})
}
