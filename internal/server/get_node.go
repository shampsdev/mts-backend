package server

import (
	"errors"
	"net/http"

	"api.mts.shamps.dev/internal/search"
	"github.com/gin-gonic/gin"
)

// NodeByID godoc
// @Summary Get a node by person id
// @Description Get a node by person id
// @Tags persons
// @Produce json
// @Param id path string true "Person ID"
// @Success 200 {object} domain.PersonNode
// @Failure 404
// @Failure 500
// @Router /persons/nodes/{id} [get]
func NodeByID(engine search.Engine) gin.HandlerFunc {
	return func(c *gin.Context) {
		personID := c.Param("id")

		node, err := engine.NodeByID(personID)

		if errors.Is(err, search.ErrNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
			return
		}

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
			return
		}

		c.JSON(http.StatusOK, node)
	}
}
