package v1

import (
	"community/config"
	"community/internal/model"
	"github.com/gin-gonic/gin"
)

func AddTag(c *gin.Context) {
	tag := model.Tag{}
	c.BindJSON(&tag)

	_, err := config.DbConn.Exec("insert into tag(name,value) "+
		"values(?,?);", tag.Name, tag.Value)
	if err != nil {
		println(err.Error())
	}
	c.JSON(200, gin.H{
		"success": 1,
		"message": "新增一条标签",
	})
}
