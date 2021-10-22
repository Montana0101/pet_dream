package api

import (
	"community/api/pet"
	"community/api/post"
	"community/api/user"
	"github.com/gin-gonic/gin"
)

func Handlers(e *gin.Engine) {
	//e := gin.New(
	e.POST("/user", user.AddUser)
	e.GET("/user/:userId", user.GetUser)

	e.POST("/pet", pet.BindPet)
	e.POST("/user/:userId/post", post.AddPost) // 发布贴文

	e.GET("/user/:userId/post/:postId", post.GetPostInfo) // 贴文详情
	e.GET("/user/:userId/posts", post.GetPostsByUser)     // 获取用户贴文

}
