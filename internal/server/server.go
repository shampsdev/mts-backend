package server

import (
	"log"

	_ "api.mts.shamps.dev/docs"
	"api.mts.shamps.dev/external/adapter"
	"api.mts.shamps.dev/internal/search"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRouter(r *gin.Engine) {
	factory := &adapter.AdapterFactory{}
	jsonAdapter, err := factory.NewAdapter(adapter.JsonAdapterType)
	if err != nil {
		log.Fatalf("failed to create json adapter: %v", err)
	}
	engine := search.NewBleveEngine(jsonAdapter)

	r.GET("/persons", GetAll(engine))
	r.GET("/persons/search", Search(engine))
	r.GET("/persons/nodes/:id", NodeByID(engine))

	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
