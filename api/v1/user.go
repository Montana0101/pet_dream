package v1

import (
	"community/config"
	"community/internal/model"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
)

func AddUser(c *gin.Context) {
	user := model.User{}
	c.BindJSON(&user)

	if user.Code == "" {
		println("未收到code")
		return
	}

	response, err := http.Get("https://api.weixin.qq.com/sns/jscode2session?appid=" + config.Appid +
		"&secret=" + config.Secret + "&js_code=" + user.Code + "&grant_type=authorization_code")

	if err != nil || response.StatusCode != http.StatusOK {
		c.Status(http.StatusServiceUnavailable)
		return
	}
	reader, err := ioutil.ReadAll(response.Body)

	var wechatLogin model.WechatLogin
	if err := json.Unmarshal([]byte(reader), &wechatLogin); err == nil {
		fmt.Println("==============json str 转map=======================")
		fmt.Println(string(reader))
		fmt.Println(wechatLogin)
	}

	if wechatLogin.Errcode == 0 && wechatLogin.Openid != "" {
		// 查找该用户是否已注册
		rows, err := config.DbConn.Query("select user.Id,user.city,user.district," +
			"user.nick_name,user.avatar_url from user where openid = ?;",
			wechatLogin.Openid)
		if err != nil {
			println(err.Error())
		}
		if rows.Next() {
			if err := rows.Scan(&user.Id, &user.City, &user.District,&user.NickName,&user.AvatarUrl); err == nil {
				c.JSON(200, gin.H{
					"success": 1,
					"message": "授权登陆成功",
					"data": gin.H{
						"user_id":  user.Id,
						"city":     user.City,
						"district": user.District,
						"openId":   wechatLogin.Openid,
						"avatar_url":user.AvatarUrl,
						"nick_name":user.NickName,
					},
				})
			}
		} else {
			//插入数据
			if res, err := config.DbConn.Exec("insert into user(openid) values(?);",
				wechatLogin.Openid); err == nil {
				userId, err := res.LastInsertId()
				if err == nil {
					c.JSON(200, gin.H{
						"success": 1,
						"message": "用户注册成功",
						"data": gin.H{
							"user_id":  userId,
							"city":     user.City,
							"district": user.District,
							"openId":   wechatLogin.Openid,
							"avatar_url":user.AvatarUrl,
							"nick_name":user.NickName,
						},
					})
				}
			}
		}
	} else {
		c.JSON(200, gin.H{
			"success": 0,
			"message": "授权登陆失败",
			"data":    wechatLogin,
		})
	}
}

// 记录地址信息
func PutLocation(c *gin.Context) {
	userId := c.Param("userId")
	user := model.User{}
	c.BindJSON(&user)
	rows, err := config.DbConn.Exec("update user set longitude=?,latitude=?,city=?,district=? "+
		"where id = ?;",
		user.Longitude, user.Latitude, user.City, user.District, userId)
	if err != nil {
		println(err.Error())
	}

	result, err := rows.RowsAffected()
	if err != nil {
		println(err.Error())
	}

	if result == 1 {
		c.JSON(200, gin.H{
			"success": 1,
			"message": "叮 ! 用户地址已变更 ≧◠◡◠≦"})
	} else {
		c.JSON(200, gin.H{
			"success": 0,
			"message": "哇 ! 用户地址变更失败（┬＿┬）"})
	}
}

// 更新微信头像和昵称
func PutUserinfo(c *gin.Context) {
	userId := c.Param("userId")
	user := model.User{}
	c.BindJSON(&user)
	rows, err := config.DbConn.Exec("update user set nick_name=?,avatar_url=? "+
		"where id = ?;",
		user.NickName, user.AvatarUrl, userId)
	if err != nil {
		println(err.Error())
	}

	result, err := rows.RowsAffected()
	if err != nil {
		println(err.Error())
	}

	if result == 1 {
		c.JSON(200, gin.H{
			"success": 1,
			"message": "叮 ! 用户微信昵称和头像已变更 ≧◠◡◠≦"})
	} else {
		c.JSON(200, gin.H{
			"success": 0,
			"message": "哇 ! 用户微信昵称和头像失败咯（┬＿┬）"})
	}
}

// 获取用户信息
func GetUser(c *gin.Context) {
	userId := c.Param("userId")
	if userId == "" {
		log.Panic("参数未传用户id")
		return
	}
	rows, err := config.DbConn.Query("select user.id,user.nick_name,user.real_name,user.age,"+
		"user.gender,user.latitude,user.longitude,user.identity,"+
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
			&user.Gender, &user.Latitude, &user.Longitude, &user.Identity,
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
			"user_id":   user.Id,
			"nick_name": user.NickName,
			"real_name": user.RealName,
			"age":       user.Age,
			"latitude":  user.Latitude,
			"longitude": user.Longitude,
			"identity":  user.Identity,
			"pet":       l,
		},
	})
}
