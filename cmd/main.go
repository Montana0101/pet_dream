package main

import (
	"community/api"
	"github.com/gin-gonic/gin"
)

func main(){
	var e *gin.Engine
	engine := gin.New()
	//engine.POST("/", api.CreateUser)
	e = engine
	api.Handlers(e)
	engine.Run() // listen and serve on 0.0.
}

//1.data model
//2.error handling.