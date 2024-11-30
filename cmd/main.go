package main

import (
	server "api.mts.shamps.dev/internal"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	server.SetupRouter(router)
	router.Run(":8000")
}
