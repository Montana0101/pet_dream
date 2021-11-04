package main

import (
	"community/api"
	"github.com/gin-gonic/gin"
)

func main() {
	var e *gin.Engine
	engine := gin.New()
	e = engine
	api.Handlers(e)
	engine.Run(":7777") // listen and serve on 0.0.
}
