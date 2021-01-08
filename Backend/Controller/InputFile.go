package Controller

import (
	"context"
	"elearn100/Pkg/e"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/qiniu/api.v7/v7/auth/qbox"
	"github.com/qiniu/api.v7/v7/storage"
	"os"
	"strings"
	"time"
)

var (
	accessKey = "qtb-RjXpvgdwd3wQbDg7cUQAm9ortDMJuW5oD0wv"
	secretKey = "JSXd2g6K1Xc8w3sK0N7MTmhgMzUPln1NcSXxg6Bx"
	bucket    = "brocas"
)

// @Summer上传图片
func UploadFile(c *gin.Context) {

	fileDir, fileErr := UploadDir()
	if !fileErr {
		e.Error(c, "目录创建失败", fileDir)
	}

	file, err := c.FormFile("file")
	if err != nil {
		e.Error(c, "没有上传图片", file)
		return
	}

	lastIndex := strings.LastIndex(file.Filename, ".")
	prefix := file.Filename[lastIndex:]
	filePath := time.Now().Format("20060102130405") + prefix
	err = c.SaveUploadedFile(file, fileDir+filePath)
	if err != nil {
		e.Error(c, "上传失败", "")
		return
	} else {
		if prefix != ".mp4" {
			filePath = QiNiu(filePath)
		}
	}

	if filePath == "" || prefix == ".mp4" {
		filePath = "/static/upload/" + time.Now().Format("20060102") + "/" + filePath
	} else {
		filePath = "http://img.cdn.brocaedu.com/" + filePath
	}
	e.Success(c, "上传成功", filePath)
}

// @Summer 创建目录
func UploadDir() (file string, isOk bool) {
	dir, _ := os.Getwd()
	upload := dir + "/Resources/Public/upload/" + time.Now().Format("20060102") + "/"
	_, err := os.Stat(upload)
	if err != nil {
		err = os.MkdirAll(upload, os.ModePerm)
		if err != nil {
			return "", false
		}
	}
	return upload, true
}

type MyPutRet struct {
	Key    string
	Hash   string
	Fsize  int
	Bucket string
	Name   string
}

func QiNiu(key string) string {
	if key == "" {
		fmt.Println("空的文件")
		return ""
	}

	fileDir, fileErr := UploadDir()
	if !fileErr {
		fmt.Println("目录创建失败")
		return ""
	}
	localFile := fileDir + key

	upToken := Ticket()
	cfg := storage.Config{}
	// 空间对应的机房
	cfg.Zone = &storage.ZoneHuanan
	// 是否使用https域名
	cfg.UseHTTPS = false
	// 上传是否使用CDN上传加速
	cfg.UseCdnDomains = false
	// 构建表单上传的对象
	formUploader := storage.NewFormUploader(&cfg)
	ret := MyPutRet{}
	// 可选配置
	putExtra := storage.PutExtra{}

	err := formUploader.PutFile(context.Background(), &ret, upToken, key, localFile, &putExtra)
	if err != nil {
		fmt.Println("七牛上传失败")
		return ""
	}
	return ret.Key
}

// @Summer 获取上传凭证
func Ticket() (upToekn string) {
	putPolicy := storage.PutPolicy{
		Scope:      bucket,
		ReturnBody: `{"key":"$(key)","hash":"$(etag)","fsize":$(fsize),"bucket":"$(bucket)","name":"$(x:name)"}`,
	}

	putPolicy.Expires = 28800

	mac := qbox.NewMac(accessKey, secretKey)
	token := putPolicy.UploadToken(mac)
	e.SetQiNiuToken(token)
	return token
}
