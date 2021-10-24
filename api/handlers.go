package api

import (
	"community/api/comment"
	"community/api/pet"
	"community/api/post"
	"community/api/user"
	"github.com/gin-gonic/gin"
)

func Handlers(e *gin.Engine) {
	//e := gin.New(
	e.POST("/user", user.AddUser)
	e.GET("/user/:userId", user.GetUser)              // 获取用户及关联的宠物信息
	e.GET("/user/:userId/posts", post.GetPostsByUser) // 获取用户贴文

	e.POST("/pet", pet.BindPet)      // 绑定宠物
	e.DELETE("/pet/:id", pet.DelPet) // 删除宠物信息
	e.PUT("/pet/:id", pet.PutPet)    // 更新宠物

	e.POST("/post", post.AddPost)                          // 发布贴文
	e.GET("/post/:id", post.GetPost)                       // 贴文详情
	e.GET("/post/:id/comments", comment.GetCommentsByPost) // 获取贴文评论

	e.POST("/comment", comment.AddComment) // 发表留言
}
