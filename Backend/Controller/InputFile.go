package Controller

import (
	"elearn100/Pkg/e"
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
	"strings"
	"time"
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
	filePath := time.Now().Format("20060102130405") + file.Filename[lastIndex:]
	err = c.SaveUploadedFile(file, fileDir+filePath)
	if err != nil {
		fmt.Print("上传失败:", err)
		e.Error(c, "上传失败", "")
		return
	}

	e.Success(c, "上传成功", time.Now().Format("20060102")+"/"+filePath)
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
