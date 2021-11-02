package api

import (
	"community/api/v1"
	"github.com/gin-gonic/gin"
)

func Handlers(e *gin.Engine) {
	e.POST("/v1/user", v1.AddUser)
	e.GET("/v1/user/:userId", v1.GetUser)              // 获取用户及关联的宠物信息
	e.GET("/v1/user/:userId/posts", v1.GetPostsByUser) // 获取用户贴文
	e.PUT("/v1/user/location", v1.PutLocation)         // 更新用户定位

	e.POST("/v1/pet", v1.BindPet)      // 绑定宠物
	e.DELETE("/v1/pet/:id", v1.DelPet) // 删除宠物信息
	e.PUT("/v1/pet/:id", v1.PutPet)    // 更新宠物

	e.POST("/v1/post", v1.AddPost)                       // 发布贴文
	e.GET("/v1/post/:id", v1.GetPost)                    // 贴文详情
	e.GET("/v1/recommend", v1.RecommendPost)             // 首页推荐贴文
	e.GET("/v1/post/:id/comments", v1.GetCommentsByPost) // 获取贴文评论
	e.POST("/v1/media", v1.SaveImg)
	e.POST("/v1/tag", v1.AddTag) // 新增标签

	e.POST("/v1/comment", v1.AddComment) // 发表留言
}
