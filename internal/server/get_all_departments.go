package server

import (
	"net/http"

	"api.mts.shamps.dev/internal/search"
	"github.com/gin-gonic/gin"
)

// GetAll godoc
// @Summary Retrieve all departments
// @Description Get a list of all departments
// @Tags persons
// @Produce json
// @Success 200 {array} string
// @Router /persons/departments [get]
func GetAllDepartments(engine search.Engine) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, engine.AllDepartments())
	}
}
