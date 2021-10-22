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
