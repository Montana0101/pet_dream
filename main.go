/*
 * @Author: your name
 * @Date: 2021-10-21 09:49:57
 * @LastEditTime: 2022-02-14 11:31:46
 * @LastEditors: Please set LastEditors
 * @Description: 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 * @FilePath: \树莓派d:\practice\cats\cmd\main.go
 */
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
	engine.Run("0.0.0.0:6122") // listen and serve on 0.0.
}
