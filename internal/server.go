package server

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Employee struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
	Role string `json:"role"`
}

var employees = []Employee{
	{ID: 1, Name: "John Doe", Role: "Software Engineer"},
	{ID: 2, Name: "Jane Smith", Role: "Product Manager"},
}

func SetupRouter(r *gin.Engine) {

	r.GET("/employees", func(c *gin.Context) {
		c.JSON(http.StatusOK, employees)
	})

	r.GET("/employees/:id", func(c *gin.Context) {
		id, _ := strconv.ParseInt(c.Param("id"), 10, 64)

		for _, employee := range employees {
			if employee.ID == id {
				c.JSON(http.StatusOK, employee)
				return
			}
		}

		c.JSON(http.StatusNotFound, gin.H{"message": "employee not found"})
	})

}
