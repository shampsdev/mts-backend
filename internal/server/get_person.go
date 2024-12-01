package server

import (
	"errors"
	"net/http"

	"api.mts.shamps.dev/internal/search"
	"github.com/gin-gonic/gin"
)

// NodeByID godoc
// @Summary Get a person by person id
// @Description Get a person by person id
// @Tags persons
// @Produce json
// @Param id path string true "Person ID"
// @Success 200 {object} domain.Person
// @Failure 404
// @Failure 500
// @Router /persons/{id} [get]
func PersonById(engine search.Engine) gin.HandlerFunc {
	return func(c *gin.Context) {
		personID := c.Param("id")

		node, err := engine.PersonById(personID)

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
