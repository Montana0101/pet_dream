package user

import (
	"community/config"
	"community/internal/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

func AddUser(c *gin.Context) {
	user := model.User{}
	c.BindJSON(&user)
	//校验昵称
	if user.NickName == "" {
		fmt.Println("添加用户失败*_*")
		return
	}
	//插入数据
	_, err := config.DbConn.Exec("insert into user(nick_name) values(?);", user.NickName)
	if err != nil {
		log.Panic(err.Error())
		return
	}
	fmt.Println("添加了一条用户数据~")
	c.JSON(200, gin.H{
		"status":  "200",
		"message": "创建用户成功",
		"data": gin.H{
			"nickname": user.NickName,
		},
	})
}

// 获取用户信息
func GetUser(c *gin.Context) {
	userId := c.Param("userId")
	if userId == "" {
		log.Panic("参数未传用户id")
		return
	}
	rows, err := config.DbConn.Query("select user.id,user.nick_name,user.real_name,user.age,"+
		"user.gender,user.latitude,user.longitude,user.identity_id,"+
		"pet.name,pet.id from user left join pet "+
		"on user.id = pet.user_id where user.id = (?) and pet.visible=1;", userId)

	defer rows.Close()

	if err != nil {
		log.Panic(err.Error())
		return
	}
	user := new(model.User)
	pet := new(model.Pet)
	var l []interface{}
	for rows.Next() {

		err := rows.Scan(&user.Id, &user.NickName, &user.RealName, &user.Age,
			&user.Gender, &user.Latitude, &user.Longitude, &user.IdentityId,
			&pet.Name, &pet.Id)
		if err != nil {
			log.Panic(err.Error())
			return
		}
		//是否有宠物
		if pet.Id != nil {
			l = append(l, gin.H{
				"name": pet.Name,
				"id":   pet.Id,
			})
		}
	}
	c.JSON(200, gin.H{
		"status":  "200",
		"message": "创建用户成功",
		"data": gin.H{
			"user_id":     user.Id,
			"nick_name":   user.NickName,
			"real_name":   user.RealName,
			"age":         user.Age,
			"latitude":    user.Latitude,
			"longitude":   user.Longitude,
			"identity_id": user.IdentityId,
			"pet":         l,
		},
	})
}
