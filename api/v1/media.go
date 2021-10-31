package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func SaveImg(c *gin.Context) {
	//xx := c.PostForm("file")
	//print("对你撒娇可能的撒娇看", xx)
	_, header, err := c.Request.FormFile("file")
	//f, h, ere := c.GetFile("file") //获取上传的文件 myfile
	//print("走打击的撒你哦大数据哦")
	if err != nil {
		print(err.Error())
	}
	//print("走大宋大撒把大家萨科")
	print(header.Size)
	//println("33333333333333333")
	pathname := header.Filename
	//println("4444444444444")
	fmt.Printf("你的撒娇可能的撒娇看", pathname)
	//fmt.Println("uuuuuuu=====", pathname)
	//fmt.Println("图片名称", pathname)
	//path := "./static/uploadout/" + pathname //图片目录+图片名称
	//fmt.Println("path===", path)
	//c.Close() //关闭上传的文件，不然的话会出现临时文件不能清除的情况
	err = c.SaveUploadedFile(header, "D:\\临时\\"+pathname)
	if err != nil {
		fmt.Printf("有四u阿达", err.Error())
	}
}
