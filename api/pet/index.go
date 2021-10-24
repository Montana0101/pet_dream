package pet

import (
	"community/config"
	"community/internal/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

func BindPet(c *gin.Context) {
	pet := model.Pet{}
	//userId := c.Param("userId")
	c.BindJSON(&pet)
	//校验必传参数
	if pet.UserId == nil || pet.Gender == nil || pet.Name == nil {
		fmt.Println("绑定宠物失败*_*")
		c.JSON(400, gin.H{
			"success": "false",
			"message": "绑定宠物失败，请核对传入参数",
		})
		return
	}

	//插入数据
	_, err := config.DbConn.Exec("insert into pet(user_id,name,age,gender) values(?,?,?,?);",
		pet.UserId, pet.Name, pet.Age, pet.Gender)
	if err != nil {
		log.Panic(err.Error())
		return
	}
	fmt.Println("叮~成功绑定了一条宠物记录≧◠◡◠≦")
	c.JSON(200, gin.H{
		"success": "200",
		"message": "绑定宠物成功",
		"data": gin.H{
			"name":   pet.Name,
			"age":    pet.Age,
			"gender": pet.Gender,
		},
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
