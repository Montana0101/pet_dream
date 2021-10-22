package post

import (
	"community/config"
	"community/internal/model"
	"github.com/gin-gonic/gin"
)

// 发布贴文
func AddPost(c *gin.Context) {
	user_id := c.Param("userId")
	post := model.Post{}
	c.BindJSON(&post)
	//校验
	if user_id == "" || post.Title == nil || post.Content == nil {
		c.JSON(400, gin.H{
			"success": 0,
			"message": "发布帖子失败，必传参数不能为空"})
		return
	}
	//新增
	_, err := config.DbConn.Exec("insert into post(user_id,title,content) "+
		"values(?,?,?);", user_id, post.Title, post.Content)
	if err != nil {
		panic(err.Error())
		return
	}
	c.JSON(200, gin.H{
		"success": 1,
		"message": "数据添加成功"})
}

// 贴文详情
func GetPostInfo(c *gin.Context) {
	post := model.Post{}
	user := model.User{}
	post_id := c.Param("postId")
	user_id := c.Param("userId")

	if post_id == "" || user_id == "" {
		println("用户或贴文id不能为空")
		return
	}

	rows, err := config.DbConn.Query("select post.id,post.title,post.content,post.user_id,user.nick_name "+
		"from post left join user on post.user_id = user.id where post.id = (?) and post.user_id=(?);",
		post_id, user_id)

	defer rows.Close()

	if err != nil {
		println("数据库查询错误", err.Error())
		return
	}

	if rows.Next() {
		err := rows.Scan(post.Id, post.Title, post.Content, post.UserId, user.NickName)
		if err == nil {
			println("数据返回错误", err.Error())
			return
		}
		c.JSON(200, gin.H{
			"success": 1,
			"message": "成功获取贴文",
			"data": gin.H{
				"id":          post_id,
				"title":       post.Title,
				"content":     post.Content,
				"user_id":     post.UserId,
				"create_time": post.CreateTime}})
	}
}

// 用户关联的贴文列表
func GetPostsByUser(c *gin.Context) {
	//user := model.User{}
	post := model.Post{}

	user_id := c.Param("user_id")
	if user_id == "" {
		println("用户Id不能为空")
		return
	}
	rows, err := config.DbConn.Query("select post.id,post.title,post.content from post where"+
		"post.user_id = (?)", user_id)
	defer rows.Close()
	if err != nil {
		println("数据库查询错误", err.Error())
		return
	}
	var list []interface{}
	for rows.Next() {
		err := rows.Scan(post.Id, post.Title, post.Content)
		if err != nil {
			println("用户Id不能为空")
			return
		}
		if post.Id != nil {
			list = append(list, gin.H{
				"id":      post.Id,
				"title":   post.Title,
				"content": post.Content,
			})
		}
	}
	c.JSON(200, gin.H{
		"success": 1,
		"message": "用户贴文列表请查收~",
		"data":    list,
	})
}
