package e

import (
	"bytes"
	"elearn100/pkg/setting"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"io/ioutil"
	"math"
	"os"
	"strconv"
	"strings"
)

// @Summary 获取绝对路径
func GetDir() string {
	dir, _ := os.Getwd()
	return dir
}

// @Summary 去除两侧空白
func Trim(con string) string {
	return strings.TrimSpace(con)
}

// @Summary 返回错误内容
func Error(ctx *gin.Context, msg string, data interface{}) {
	ctx.JSON(SUCCESS, gin.H{
		"code": ERROR,
		"msg":  msg,
		"data": data,
	})
}

// @Summary 返回正确内容
func Success(ctx *gin.Context, msg string, data interface{}) {
	ctx.SecureJSON(SUCCESS, gin.H{
		"code": SUCCESS,
		"msg":  msg,
		"data": data,
	})
}

func SendRes(ctx *gin.Context, code int, msg string, data interface{}) {
	ctx.SecureJSON(SUCCESS, gin.H{
		"code": code,
		"msg":  msg,
		"data": data,
	})
}

func GetBody(c *gin.Context) io.ReadCloser {
	buf := make([]byte, 1024)
	n, _ := c.Request.Body.Read(buf)
	return ioutil.NopCloser(bytes.NewReader(buf[:n]))
}

// @Summer 返回可以分页的总数
// @Param pageNum int 分页总数
func GetPageNum(count int) float64 {
	pageNum := math.Ceil(float64(count) / float64(setting.PageSize))
	return pageNum
}

func SendMessage(code int, msg string) map[string]interface{} {
	data := make(map[string]interface{})
	data["code"] = code
	data["msg"] = msg
	return data
}

func GetUUID(c *gin.Context) (isOK bool, UUID int) {
	uuid, err := c.Cookie("uuid")
	if err != nil {
		fmt.Println("没有拿到uid")
		return false, 0
	}
	uid, _ := strconv.Atoi(uuid)
	return true, uid
}

func SetCookie(c *gin.Context, uuid, maxAge int) {
	c.SetCookie("uuid", strconv.Itoa(uuid), maxAge, "/", setting.Domain, false, true)
}
