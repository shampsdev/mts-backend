package server

import (
	"net/http"

	"api.mts.shamps.dev/internal/search"
	"github.com/gin-gonic/gin"
)

func Search(engine search.Engine) gin.HandlerFunc {
	return func(c *gin.Context) {
		text := c.Query("text")
		if text == "" {
			c.JSON(http.StatusBadRequest, gin.H{"message": "search query is empty"})
			return
		}

		c.JSON(http.StatusOK, engine.Search(text))
	}
}
