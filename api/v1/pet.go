package v1

import (
	"community/config"
	"community/internal/enum"
	"community/internal/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

const (
//	　　 美国短毛猫 短
//usShorthair = iota
//ukShorthair
)

func BindPet(c *gin.Context) {
	print(enum.UkShorthair)
	pet := model.Pet{}
	c.BindJSON(&pet)
	//校验必传参数
	if pet.UserId == nil || pet.Gender == nil || pet.Name == nil || pet.Breed == nil {
		fmt.Println("绑定宠物失败*_*")
		c.JSON(200, gin.H{
			"success": 0,
			"message": "绑定宠物失败，请核对传入参数",
		})
		return
	}

	//插入数据
	_, err := config.DbConn.Exec("insert into pet(user_id,name,age,gender,breed,intro) values(?,?,?,?,?,?);",
		pet.UserId, pet.Name, pet.Age, pet.Gender, pet.Breed, pet.Intro)
	if err != nil {
		log.Panic(err.Error())
		return
	}
	fmt.Println("叮~成功绑定了一条宠物记录≧◠◡◠≦")
	c.JSON(200, gin.H{
		"success": 1,
		"message": "绑定宠物成功",
		"data": gin.H{
			"name":   pet.Name,
			"age":    pet.Age,
			"gender": pet.Gender,
		},
	})
}

// 用户关联宠物列表
func BoundPets(c *gin.Context) {
	userId := c.Param("id")
	pet := model.Pet{}
	rows, err := config.DbConn.Query("select id,name from pet where user_id = ?", userId)
	if err != nil {
		println(err.Error())
		return
	}
	var list []interface{}

	for rows.Next() {
		if err := rows.Scan(&pet.Id, &pet.Name); err == nil {
			list = append(list, gin.H{
				"id":   pet.Id,
				"name": pet.Name,
			})
		}
	}
	c.JSON(200, gin.H{
		"success": 1,
		"message": "返回用户关联宠物信息成功",
		"data":    list,
	})
}

// 删除宠物记录
func DelPet(c *gin.Context) {
	//userId := c.Param("userId")
	id := c.Param("id")
	//_, err := config.DbConn.Exec("delete from pet where user_id = (?) and id = (?);", userId, id)
	rows, err := config.DbConn.Exec("update pet set visible = 0 where id = (?) ;", id)
	if err != nil {
		fmt.Print(err.Error())
	}
	result, err := rows.RowsAffected()
	if err != nil {
		fmt.Print(err.Error())
	}
	if result == 1 {
		c.JSON(200, gin.H{
			"success": 1,
			"message": "叮 ! 成功解除与宠物关系 ≧◠◡◠≦"})
	} else {
		c.JSON(200, gin.H{
			"success": 0,
			"message": "哇 ! 解除与宠物关系失败（┬＿┬）"})
	}
}

// 更新宠物信息
func PutPet(c *gin.Context) {
	pet := model.Pet{}
	id := c.Param("id")
	c.BindJSON(&pet)
	rows, err := config.DbConn.Exec("update pet set name=?,age=?,gender=? where id = ?;",
		pet.Name, pet.Age, pet.Gender, id)
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
			"message": "叮 ! 宠物信息已变更 ≧◠◡◠≦"})
	} else {
		c.JSON(200, gin.H{
			"success": 0,
			"message": "哇 ! 信息变更失败（┬＿┬）"})
	}
}
