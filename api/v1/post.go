package v1

import (
	"community/config"
	"community/internal/model"
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

// AddPost 发布贴文
func AddPost(c *gin.Context) {
	//user_id := c.Param("userId")
	post := model.Post{}
	c.BindJSON(&post)

	//校验
	if post.UserId == nil || post.PetId == nil || post.Title == nil || post.Content == nil {
		c.JSON(400, gin.H{
			"success": 0,
			"message": "发布帖子失败，必传参数不能为空"})
		return
	}

	//新增
	res, err := config.DbConn.Exec("insert into post(user_id,pet_id,title,content) "+
		"values(?,?,?,?);", post.UserId, post.PetId, post.Title, post.Content)
	if err != nil {
		println(err.Error())
	}
	// 返回自增id
	postId, err := res.LastInsertId()
	print("返回自增id", postId)

	// 判断有资源
	if len(post.Media) > 0 {
		for i, _ := range post.Media {
			//存照片视频
			if _, err := config.DbConn.Exec("insert into media(url,type,user_id,post_id) values(?,?,?,?)",
				post.Media[i].Url, post.Media[i].Type, post.UserId, postId); err == nil {
				print("返回自增dsadsadsid", postId)
				println("遍历资源到数据库成功咯")
			} else {
				println("遍历资源到数据库失败拉", err.Error())
				c.JSON(200, gin.H{
					"success": 0,
					"message": "图片添加失败"})
				return
			}
		}
	}

	c.JSON(200, gin.H{
		"success": 1,
		"message": "数据添加成功",
		"data": gin.H{
			"id": postId,
		}})
}

// GetPost 贴文详情
func GetPost(c *gin.Context) {
	post := model.Post{}
	user := model.User{}
	pet := model.Pet{}
	postId := c.Param("id")

	if postId == "" {
		println("用户或贴文id不能为空")
		return
	}

	rows, err := config.DbConn.Query("select post.id,post.title,post.content,post.user_id,"+
		"post.create_time,user.nick_name,user.city,user.avatar_url,pet.name,pet.age,pet.gender,"+
		"pet.breed,pet.photo "+
		"from post left join user on post.user_id = user.id left join pet on post.pet_id=pet.id "+
		"where post.id = (?)", postId)

	if err != nil {
		println("数据库查询贴文错误", err.Error())
		return
	}

	imgs, err := config.DbConn.Query("select url,type from media where post_id = ?", postId)
	if err != nil {
		println("数据库查询图片错误", err.Error())
		return
	}

	var arr []interface{}
	media := model.Media{}
	for imgs.Next() {
		err := imgs.Scan(&media.Url, &media.Type)
		if err != nil {
			println(err.Error())
		}
		arr = append(arr, gin.H{
			"url":  media.Url,
			"type": media.Type,
		})
	}

	if rows.Next() {
		err := rows.Scan(&post.Id, &post.Title, &post.Content, &post.UserId, &post.CreateTime,
			&user.NickName, &user.City, &user.AvatarUrl, &pet.Name, &pet.Age, &pet.Gender,
			&pet.Breed, &pet.Photo)
		if err != nil {
			println("数据返回错误", err.Error())
			return
		}
		c.JSON(200, gin.H{
			"success": 1,
			"message": "成功获取贴文",
			"data": gin.H{
				"id":         postId,
				"title":      post.Title,
				"content":    post.Content,
				"userId":     post.UserId,
				"nickName":   user.NickName,
				"city":       user.City,
				"avatarUrl":  user.AvatarUrl,
				"createTime": post.CreateTime,
				"petInfo": gin.H{
					"name":   pet.Name,
					"gender": pet.Gender,
					"breed":  pet.Breed,
					"age":    pet.Age,
					"photo":  pet.Photo,
				},
				"images": arr}})
	} else {
		c.JSON(200, gin.H{
			"success": 0,
			"message": "未查到该数据"})
	}
}

// 用户关联的贴文列表
func GetPostsByUser(c *gin.Context) {
	post := model.Post{}
	userId := c.Param("userId")

	if userId == "" {
		println("用户Id不能为空")
		return
	} else {
		println("用户有id")
	}

	rows, err := config.DbConn.Query("select post.id,post.title,post.content from post where post.user_id=(?);", userId)

	if err != nil {
		println("数据库查询错误", err.Error())
	}

	var list []interface{}

	for rows.Next() {
		err := rows.Scan(&post.Id, &post.Title, &post.Content)
		if err != nil {
			println(err.Error())
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

// 首页推荐推文
func RecommendPost(c *gin.Context) {
	post := model.Post{}
	user := model.User{}
	media := model.Media{}
	pet := model.Pet{}
	//district := c.Param("district")

	// 分页查询
	pageNo := c.DefaultQuery("pageNo", "1")
	pageSize := c.DefaultQuery("pageSize", "10")
	gender := c.DefaultQuery("gender", "")
	city := c.DefaultQuery("city", "")
	// 字符串转数字
	no, _ := strconv.Atoi(pageNo)
	size, _ := strconv.Atoi(pageSize)
	cutNo := (no - 1) * size

	var err error
	//var rows interface{}
	//rows := sql.Rows{}
	var rows *sql.Rows
	if city == "" {
		var sql = "select user.id as user_id,user.nick_name,user.city," +
			"user.district,post.id as post_id,post.title,post.content,media.url,media.type," +
			"pet.name,pet.age,pet.gender from post " +
			"left join user on post.user_id = user.id left join (select type,id,url,post_id from media group by post_id)" +
			"as media on post.id = media.post_id " +
			"left join pet on post.pet_id = pet.id " +
			"where 1=1 %s order by post.create_time desc limit ?,?;"
		if gender != "" {
			//sql = fmt.Sprintf(sql)
			sql = fmt.Sprintf(sql, "and pet.gender="+gender)
			println("打印下sql", sql)
		} else {
			sql = fmt.Sprintf(sql, "and 2=2")
		}

		// 查全国
		rows, err = config.DbConn.Query(sql, cutNo, pageSize)
	} else {
		// 查所在城市
		rows, err = config.DbConn.Query("select user.id as user_id,user.nick_name,user.city,"+
			"user.district,post.id as post_id,post.title,post.content,media.url,media.type,"+
			"pet.name,pet.age,pet.gender from post "+
			"inner join user on post.user_id = user.id and user.city=?"+
			"left join (select type,id,url,post_id from media group by post_id)"+
			"as media on post.id = media.post_id "+
			"left join pet on post.pet_id=pet.id order by post.create_time desc limit ?,?;", city, cutNo, pageSize)
	}
	if err != nil {
		println("数据库查询错误", err.Error())

	}
	var list []interface{}
	count := 0

	for rows.Next() {
		count++
		if err := rows.Scan(&user.Id, &user.NickName, &user.City, &user.District,
			&post.Id, &post.Title, &post.Content, &media.Url, &media.Type,
			&pet.Name, &pet.Age, &pet.Gender); err == nil {
			if post.Id != nil {
				list = append(list, gin.H{
					"userId":    user.Id,
					"nickname":  user.NickName,
					"city":      user.City,
					"district":  user.District,
					"postId":    post.Id,
					"title":     post.Title,
					"content":   post.Content,
					"mediaUrl":  media.Url,
					"mediaType": media.Type,
					"petName":   pet.Name,
					"petAge":    pet.Age,
					"petGender": pet.Gender,
				})
			}
		}
	}

	c.JSON(200, gin.H{
		"success": 1,
		"message": city + "的贴文列表请查收~",
		"data":    list,
		"count":   count,
	})
}
