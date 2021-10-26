package v1

import (
	"fmt"
	"io"
	"strconv"
	"time"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

//func OssConnect() {
//	// 创建OSSClient实例。
//	// yourEndpoint填写Bucket所在地域对应的Endpoint。以华东1（杭州）为例，Endpoint填写为https://oss-cn-hangzhou.aliyuncs.com。
//	// 阿里云账号AccessKey拥有所有API的访问权限，风险很高。强烈建议您创建并使用RAM用户进行API访问或日常运维，请登录RAM控制台创建RAM用户。
//	client, err := oss.New("https://oss-cn-shanghai.aliyuncs.com", "LTAI5tJCrERAYLx4iGacjepD", "eBOHNnw81KVtZ3cVNQSBe4l8rvF6jk")
//	if err != nil {
//		fmt.Println("Error:", err)
//		os.Exit(-1)
//	}
//	// 填写Bucket名称。
//	bucketName := " mini-weather"
//	// 填写Object完整路径。Object完整路径中不能包含Bucket名称。
//	objectName := "exampleobject.txt"
//	// 填写本地文件的完整路径。如果未指定本地路径，则默认从示例程序所属项目对应本地路径中上传文件。
//	locaFilename := "D:\\localpath\\examplefile.txt"
//
//	// 获取存储空间。
//	bucket, err := client.Bucket(bucketName)
//	if err != nil {
//		fmt.Println("Error:", err)
//		os.Exit(-1)
//	}
//	// 将本地文件分片，且分片数量指定为3。
//	chunks, err := oss.SplitFileByPartNum(locaFilename, 3)
//	fd, err := os.Open(locaFilename)
//	defer fd.Close()
//
//	// 设置存储类型为标准存储。
//	storageType := oss.ObjectStorageClass(oss.StorageStandard)
//
//	// 步骤1：初始化一个分片上传事件，并指定存储类型为标准存储。
//	imur, err := bucket.InitiateMultipartUpload(objectName, storageType)
//	// 步骤2：上传分片。
//	var parts []oss.UploadPart
//	for _, chunk := range chunks {
//		fd.Seek(chunk.Offset, os.SEEK_SET)
//		// 调用UploadPart方法上传每个分片。
//		part, err := bucket.UploadPart(imur, fd, chunk.Size, chunk.Number)
//		if err != nil {
//			fmt.Println("Error:", err)
//			os.Exit(-1)
//		}
//		parts = append(parts, part)
//	}
//
//	// 指定Object的读写权限为公共读，默认为继承Bucket的读写权限。
//	objectAcl := oss.ObjectACL(oss.ACLPublicRead)
//
//	// 步骤3：完成分片上传，指定文件读写权限为公共读。
//	cmur, err := bucket.CompleteMultipartUpload(imur, parts, objectAcl)
//	if err != nil {
//		fmt.Println("Error:", err)
//		os.Exit(-1)
//	}
//	fmt.Println("cmur:", cmur)
//}

type Config struct {
	Endpoint        string
	AccessKeyId     string
	AccessKeySecret string
	Region          string
	Bucket          string
	Secure          bool
	Cname           bool
}

func (this *Config) OssCennect() (Bucket *oss.Bucket, err error) {
	//此处需要进入阿里云oss控制台配置域名
	this.Endpoint = "https://image.xxxxxxx.com"
	this.AccessKeyId = "你的accessKeyId"
	this.AccessKeySecret = "你的Secret"
	this.Region = "选区:我的oss-cn-beijing"
	this.Bucket = "bucket名称"
	this.Secure = true
	this.Cname = true
	client, err := oss.New(this.Endpoint, this.AccessKeyId, this.AccessKeySecret, oss.UseCname(true))
	if err != nil {
		fmt.Println(err)
	}
	Bucket, err = client.Bucket("tybk")

	return

}

func (this *Config) LocalUrl(file io.Reader) (url string, err error) {
	Bucket, err := this.OssCennect()
	if err != nil {
		panic(err)
	}

	t := time.Now()
	fmt.Println(t.Year())
	//拼接文件名称
	fileName := fmt.Sprintf("%s%s%s%s%s%s%s", strconv.Itoa(t.Year()), strconv.Itoa(int(t.Month())), strconv.Itoa(t.Day()), strconv.Itoa(t.Hour()), strconv.Itoa(t.Minute()), strconv.Itoa(t.Second()), strconv.Itoa(int(t.Unix())))
	fmt.Println("...s......................")
	fmt.Println(fileName)
	str := "test/" + fileName + ".jpg"

	err = Bucket.PutObject(str, file)
	if err != nil {
		url = "上传错误"
	} else {
		url = fmt.Sprintf("%s%s", "https://static.xxxxxxxi.cn/", str)
	}
	return url, err

}
