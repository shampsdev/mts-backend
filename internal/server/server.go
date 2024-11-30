package server

import (
	"api.mts.shamps.dev/internal/search"
	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine) {
	engine := search.NewJSONEngine()

	r.GET("/persons", GetAll(engine))
	r.GET("/persons/filter", Filter(engine))
	r.GET("/persons/search", Search(engine))
}
