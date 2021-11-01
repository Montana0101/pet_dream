package v1

import (
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/gin-gonic/gin"
	"os"
)

type OSSConfig struct {
	Endpoint        string
	AccessKeyId     string
	AccessKeySecret string
	Region          string
	Bucket          string
	Secure          bool
	Cname           bool
}

func OssCennect() (Bucket *oss.Bucket, err error) {
	config := OSSConfig{}
	//此处需要进入阿里云oss控制台配置域名
	config.Endpoint = "oss-cn-shanghai.aliyuncs.com"
	config.AccessKeyId = "LTAI5tJCrERAYLx4iGacjepD"
	config.AccessKeySecret = "eBOHNnw81KVtZ3cVNQSBe4l8rvF6jk"
	//config.Region = "oss-cn-shanghai"
	config.Bucket = "mini-weather"
	//config.Secure = true
	//config.Cname = true
	client, err := oss.New(config.Endpoint, config.AccessKeyId, config.AccessKeySecret)
	if err != nil {
		fmt.Println(err)
	}
	Bucket, err = client.Bucket(config.Bucket)
	if err != nil {
		fmt.Printf("链接Bucket出错啦", err.Error())
		os.Exit(-1)
	}
	return
}

func SaveImg(c *gin.Context) {
	file, header, err := c.Request.FormFile("file")
	Bucket, err := OssCennect()
	config := OSSConfig{}
	config.Bucket = "mini-weather"
	config.Endpoint = "oss-cn-shanghai.aliyuncs.com"
	if err != nil {
		print(err.Error())
	}
	pathname := header.Filename
	fmt.Printf("答应下名称", pathname)
	str := "pet/" + pathname
	print(file)
	if err = Bucket.PutObject(str, file); err != nil {
		fmt.Printf("打印下错误", err)
		c.JSON(200, gin.H{
			"success": 0,
			"message": "上传失败",
		})
	} else {
		c.JSON(200, gin.H{
			"success": 1,
			"message": "上传成功",
			"url":     "https://" + config.Bucket + "." + config.Endpoint + "/" + str,
		})
	}
}
