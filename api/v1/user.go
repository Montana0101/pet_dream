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
		rows, err := config.DbConn.Query("select user.city,user.district from user where openid = ?;",
			wechatLogin.Openid)
		if err != nil {
			println(err.Error())
		}
		if rows.Next() {
			if err := rows.Scan(&user.City, &user.District); err == nil {
				c.JSON(200, gin.H{
					"success": 1,
					"message": "授权登陆成功",
					"data": gin.H{
						"city":     user.City,
						"district": user.District,
						"openId":   wechatLogin.Openid,
					},
				})
			}
		} else {
			//插入数据
			if _, err := config.DbConn.Exec("insert into user(openid) values(?);",
				wechatLogin.Openid); err == nil {
				c.JSON(200, gin.H{
					"success": 1,
					"message": "用户注册成功",
					"data": gin.H{
						"city":     user.City,
						"district": user.District,
						"openId":   wechatLogin.Openid,
					},
				})
			} else {
				print(err.Error())
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
	user := model.User{}
	c.BindJSON(&user)
	if user.Longitude == nil || user.Latitude == nil || user.City == nil || user.Openid == nil {
		println("缺少必要参数")
		return
	}
	rows, err := config.DbConn.Exec("update user set longitude=?,latitude=?,city=? "+
		"where openid = ?;",
		user.Longitude, user.Latitude, user.City, user.Openid)
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
