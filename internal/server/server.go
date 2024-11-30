package server

import (
	_ "api.mts.shamps.dev/docs"
	"api.mts.shamps.dev/internal/search"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRouter(r *gin.Engine) {
	engine := search.NewJSONEngine()

	r.GET("/persons", GetAll(engine))
	r.GET("/persons/filter", Filter(engine))
	r.GET("/persons/search", Search(engine))

	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
