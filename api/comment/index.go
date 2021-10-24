package comment

import (
	"community/config"
	"community/internal/model"
	"github.com/gin-gonic/gin"
)

// AddComment 发布贴文
func AddComment(c *gin.Context) {
	comment := model.Comment{}
	c.BindJSON(&comment)
	if comment.UserId == nil || comment.PostId == nil || comment.Content == nil || comment.Source == nil {
		println("发布留言失败")
		return
	}
	_, err := config.DbConn.Exec("insert into comment(user_id,post_id,content,source) "+
		"values(?,?,?,?);", comment.UserId, comment.PostId, comment.Content, comment.Source)
	if err != nil {
		c.JSON(400, gin.H{
			"success": 0,
			"message": "留言发布失败，必传参数不能为空"})
	}
	c.JSON(200, gin.H{
		"success": 1,
		"message": "留言发布成功"})
}

// 获取评论集
func GetCommentsByPost(c *gin.Context) {
	comment := model.Comment{}
	postId := c.Param("id")
	rows, err := config.DbConn.Query("select comment.id,comment.content,comment.create_time "+
		"from comment where comment.post_id = ?;", postId)
	//.Scan(&comment.Id, &comment.Content,&comment.CreateTime)
	if err != nil {
		println(err.Error())
	}
	var list []interface{}

	for rows.Next() {
		err := rows.Scan(&comment.Id, &comment.Content, &comment.CreateTime)
		if err != nil {
			println(err.Error())
		}
		if comment.Id != nil {
			list = append(list, gin.H{
				"id":          comment.Id,
				"content":     comment.Content,
				"create_time": comment.CreateTime})
		}
	}
	c.JSON(200, gin.H{
		"success": 1,
		"message": "贴文评论集合请查收~",
		"data":    list,
	})
}
