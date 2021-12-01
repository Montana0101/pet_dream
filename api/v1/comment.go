package v1

import (
	"community/config"
	"community/internal/model"
	"fmt"
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
		if _, err := config.DbConn.Exec("insert into comment(user_id,post_id,parent_id,reply_id,content,source) "+
			"values(?,?,?,?,?,?);", comment.UserId, comment.PostId, comment.ParentId, comment.ReplyId, comment.Content, comment.Source); err != nil {
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
	rows, err := config.DbConn.Query("select comment.id,comment.reply_id,comment.content,"+
		"comment.reply_id,IFNULL(comment.parent_id,0),comment.source,comment.create_time, "+
		"user.nick_name,user.avatar_url from comment left join user on comment.user_id = user.id "+
		"where comment.post_id = ? and user.avatar_url is not null;", postId)
	if err != nil {
		println(err.Error())
	}

	var list []interface{}

	for rows.Next() {
		err := rows.Scan(&comment.Id, &comment.ReplyId, &comment.Content, &comment.ReplyId,
			&comment.ParentId, &comment.Source,
			&comment.CreateTime, &user.NickName, &user.AvatarUrl)

		if err != nil {
			print(111111111111111)
			println(err.Error())
			return
		}

		if comment.Id != nil {
			list = append(list, map[string]interface{}{
				"id":          comment.Id,
				"content":     comment.Content,
				"source":      comment.Source,
				"create_time": comment.CreateTime,
				"nick_name":   user.NickName,
				"avatar_url":  user.AvatarUrl,
				"parent_id":   comment.ParentId,
				"reply_id":    comment.ReplyId,
			})
		}
	}

	var list1 []interface{}
	var list2 []interface{}

	for i := 0; i < len(list); i++ {
		//m := make(list[i] interface {})
		res, _ := list[i].(map[string]interface{})

		if res["parent_id"] == 0 {
			list1 = append(list1, gin.H{
				"id":          res["id"],
				"content":     res["content"],
				"source":      res["source"],
				"create_time": res["create_time"],
				"nick_name":   res["nick_name"],
				"avatar_url":  res["avatar_url"],
				"parent_id":   res["parent_id"],
				"reply_id":    res["reply_id"],
			})
		} else {

			list2 = append(list2, gin.H{
				"id":          res["id"],
				"content":     res["content"],
				"source":      res["source"],
				"create_time": res["create_time"],
				"nick_name":   res["nick_name"],
				"avatar_url":  res["avatar_url"],
				"parent_id":   res["parent_id"],
				"reply_id":    res["reply_id"],
			})
		}
	}

	for i, _ := range list1 {
		r1, _ := list1[i].(map[string]interface{})
		var arr []interface{}

		for i2, _ := range list2 {
			r2, _ := list2[i2].(map[string]interface{})

			if r2["parent_id"] == r1["id"] {
				fmt.Printf("%v \n", "第三代就看撒")
				arr = append(arr, gin.H{
					"asd": "的撒",
				})
			}
		}
		//r1["xx"] = make(map[string]string)
		//r1["xx"] = "但是那就快点撒"
		//r1 := make(map[string]map[int]int)

		fmt.Printf("%v \n", r1)
	}

	c.JSON(200, gin.H{
		"success": 1,
		"message": "贴文评论集合请查收~",
		"data":    list1,
	})

}
