package v1

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

	if _, err := config.DbConn.Exec("insert into comment(user_id,post_id,content,source) values(?,?,?,?);",
		comment.UserId, comment.PostId, comment.Content, comment.Source); err != nil {
		println("错撒旦把数据库")
		println(err.Error())
		return
	}

	c.JSON(200, gin.H{
		"success": 1,
		"message": "留言发布成功"})
}

// 回复评论
func ReplyComment(c *gin.Context) {
	comment := model.Comment{}
	c.BindJSON(&comment)
	if comment.UserId == nil || comment.PostId == nil || comment.Content == nil || comment.Source == nil || comment.ParentId == 0 {
		println("缺少必要参数")
		return
	}

	if comment.ReplyId == nil {
		if _, err := config.DbConn.Exec("insert into comment(user_id,post_id,parent_id,content,source) "+
			"values(?,?,?,?,?);", comment.UserId, comment.PostId, comment.ParentId, comment.Content, comment.Source); err != nil {
			println(err.Error())
			return
		}
	} else {
		if _, err := config.DbConn.Exec("insert into comment(user_id,post_id,parent_id,reply_id,reply_name,content,source) "+
			"values(?,?,?,?,?,?,?);", comment.UserId, comment.PostId, comment.ParentId, comment.ReplyId,
			comment.ReplyName, comment.Content, comment.Source); err != nil {
			println(err.Error())
			return
		}
	}

	c.JSON(200, gin.H{
		"success": 1,
		"message": "留言发布成功"})
}

// 获取评论集
func GetCommentsByPost(c *gin.Context) {
	comment := model.Comment{}
	user := model.User{}
	postId := c.Param("id")
	rows, err := config.DbConn.Query("select comment.id,comment.reply_id,comment.reply_name,comment.content,"+
		"IFNULL(comment.parent_id,0),comment.source,comment.create_time, "+
		"user.nick_name,user.avatar_url from comment left join user on comment.user_id = user.id "+
		"where comment.post_id = ? and user.avatar_url is not null;", postId)

	if err != nil {
		println(err.Error())
	}

	var list []interface{}

	for rows.Next() {
		err := rows.Scan(&comment.Id, &comment.ReplyId, &comment.ReplyName, &comment.Content,
			&comment.ParentId, &comment.Source,
			&comment.CreateTime, &user.NickName, &user.AvatarUrl)

		if err != nil {
			println(err.Error())
			return
		}

		if comment.Id != 0 {
			list = append(list, map[string]interface{}{
				"id":          comment.Id,
				"content":     comment.Content,
				"source":      comment.Source,
				"create_time": comment.CreateTime,
				"nick_name":   user.NickName,
				"avatar_url":  user.AvatarUrl,
				"parent_id":   comment.ParentId,
				"reply_id":    comment.ReplyId,
				"reply_name":  comment.ReplyName,
			})
		}
	}

	var list1 []interface{}
	var list2 []interface{}

	for i := 0; i < len(list); i++ {
		//m := make(list[i] interface {})
		res, _ := list[i].(map[string]interface{})

		if res["parent_id"] == 0 {
			list1 = append(list1, map[string]interface{}{
				"id":          res["id"],
				"content":     res["content"],
				"source":      res["source"],
				"create_time": res["create_time"],
				"nick_name":   res["nick_name"],
				"avatar_url":  res["avatar_url"],
				"reply_id":    res["reply_id"],
				"reply_name":  res["reply_name"],
				"data":        []interface{}{},
			})
		} else {
			list2 = append(list2, map[string]interface{}{
				"id":          res["id"],
				"content":     res["content"],
				"source":      res["source"],
				"create_time": res["create_time"],
				"nick_name":   res["nick_name"],
				"avatar_url":  res["avatar_url"],
				"parent_id":   res["parent_id"],
				"reply_id":    res["reply_id"],
				"reply_name":  res["reply_name"],
			})
		}
	}

	for i, _ := range list1 {
		id := list1[i].(map[string]interface{})["id"]
		var data []interface{}
		if len(list2) > 0 {
			for i2, _ := range list2 {
				parentId := list2[i2].(map[string]interface{})["parent_id"]

				if parentId == id {
					data = append(data, list2[i2])
				}
			}
		}
		list1[i].(map[string]interface{})["data"] = data
	}

	c.JSON(200, gin.H{
		"success": 1,
		"message": "贴文评论集合请查收~",
		"data":    list1,
	})

}
