package server

import (
	"net/http"

	"api.mts.shamps.dev/internal/search"
	"github.com/gin-gonic/gin"
)

// FindPath godoc
// @Summary Find path between two persons
// @Description Find connection path between two persons by IDs
// @Tags persons
// @Produce json
// @Param from query string true "ID of the starting person"
// @Param to query string true "ID of the target person"
// @Success 200 {array} domain.PersonNode
// @Failure 400
// @Failure 404
// @Router /persons/nodes/path [get]
func FindPath(engine search.Engine) gin.HandlerFunc {
	return func(c *gin.Context) {
		from := c.Query("from")
		to := c.Query("to")

		// Проверка наличия параметров
		if from == "" || to == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "missing or invalid parameters"})
			return
		}

		path, err := engine.FindPathByIDs(from, to)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, path)
	}
}
