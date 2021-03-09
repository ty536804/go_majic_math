package e

import (
	"bytes"
	"elearn100/Pkg/setting"
	"fmt"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"io"
	"io/ioutil"
	"math"
	"os"
	"regexp"
	"strings"
	"time"
)

const (
	REDISKey  = "eLearn:"
	VALIDTime = "86400*30"
	MENUKey   = REDISKey + "menu:show" //导航key
	Token     = REDISKey + "token"
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

// @Desc 手机号验证
func CheckPhone(tel string) bool {
	reg := regexp.MustCompile(`^1{1}\d{10}$`)
	if !reg.MatchString(tel) || len(tel) < 11 {
		return false
	}
	return true
}

// 去除html标签
func TrimHtml(src string) string {
	//将HTML标签全转换成小写
	re, _ := regexp.Compile("\\<[\\S\\s]+?\\>")
	src = re.ReplaceAllStringFunc(src, strings.ToLower)
	//去除STYLE
	re, _ = regexp.Compile("\\<style[\\S\\s]+?\\</style\\>")
	src = re.ReplaceAllString(src, "")
	//去除SCRIPT
	re, _ = regexp.Compile("\\<script[\\S\\s]+?\\</script\\>")
	src = re.ReplaceAllString(src, "")
	//去除所有尖括号内的HTML代码，并换成换行符
	re, _ = regexp.Compile("\\<[\\S\\s]+?\\>")
	src = re.ReplaceAllString(src, "\n")
	//去除连续的换行符
	re, _ = regexp.Compile("\\s{2,}")
	src = re.ReplaceAllString(src, "\n")
	return strings.TrimSpace(src)
}

// @Desc 解析错误原因
func ViewErr(valid validation.Validation) (code int, err string) {
	for _, err := range valid.Errors {
		return ERROR, err.Message
	}
	return ReSuccess()
}

// @Desc 返回正确信息
func ReSuccess() (code int, err string) {
	return SUCCESS, "操作成功"
}

// @Desc 返回错误信息
func ReError() (int, string) {
	return SUCCESS, "操作失败"
}

// @Desc 号码不能为空
func ValidTel(tel, land, client string) bool {
	if tel == "" && land == "" && client == "" {
		return true
	}
	return false
}

func SubUUID(RemoteAddr string) string {
	return strings.Split(strings.Replace(RemoteAddr, ".", "", -1), ":")[0]
}

func GetFirstUrl(Referer, host, url, reqURI string) string {
	if Referer == "" {
		return setting.ReplaceSiteUrl(host, url, reqURI) //来源页
	}
	return Referer
}

// @Desc 去除IP地址后面的端口号
func GetIpAddress(ip string) string {
	if ipIndex := strings.LastIndex(ip, ":"); ipIndex != -1 {
		return ip[0:ipIndex]
	}
	return ip
}

// @Desc 创建日志目录
func CreateLogDir() string {
	dir := GetDir()
	dir = dir + "/Log/" + time.Now().Format("20060102") + "/"
	_, err := os.Stat(dir)
	if err != nil {
		err := os.Mkdir(dir, os.ModePerm)
		if err != nil {
			fmt.Printf("目录创建失败:%s \n", err)
		}
		return dir
	}
	return dir
}

func GetCurrentTime() string {
	return time.Now().Format("2006-01-02 15:04:05")
}
