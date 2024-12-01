package server

import (
	"net/http"

	"api.mts.shamps.dev/internal/search"
	"github.com/gin-gonic/gin"
)

// Search godoc
// @Summary Search for persons
// @Description Search for persons using a text query
// @Tags persons
// @Produce json
// @Param text query string true "Text to search for"
// @Success 200 {array} domain.Person
// @Failure 400 {object} map[string]string "Bad Request"
// @Router /persons/search [get]
func Search(engine search.Engine) gin.HandlerFunc {
	return func(c *gin.Context) {
		text := c.Query("text")
		if text == "" {
			c.JSON(http.StatusBadRequest, gin.H{"message": "search query is empty"})
			return
		}

		persons, err := engine.SearchPersons(text, []search.Filter{})
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		c.JSON(http.StatusOK, persons)
	}
}
