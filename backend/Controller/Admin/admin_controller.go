package Admin

import (
	"elearn100/Model/Admin"
	"elearn100/Services"
	"elearn100/pkg/e"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"net/http"
	"runtime"
	"time"
)

// @Summer 管理员登录
func Login(c *gin.Context) {
	code, msg := Services.Login(c)
	e.SendRes(c, code, msg, "")
}

// @Summer 后端首页
func Show(c *gin.Context) {
	userInfo := Services.GetUserById(c)
	user := Admin.SysAdminUser{}
	json.Unmarshal([]byte(userInfo), &user)
	c.HTML(e.SUCCESS, "admin/home.html", gin.H{
		"title":      "易学教育",
		"user":       user,
		"target_url": "/api/v1/index",
	})
}

// @Summer 后端首页详情内容
func BackEndIndex(c *gin.Context) {
	ginVersion := gin.Version
	osVersion := runtime.Version()
	os := runtime.GOOS
	currentTime := time.Now().Format("2006:01:02 15:04:05")

	userInfo := Services.GetUserById(c)
	user := Admin.SysAdminUser{}
	json.Unmarshal([]byte(userInfo), &user)
	fmt.Println("首页详情内容", user.NickName)
	c.HTML(e.SUCCESS, "admin/welcome.html", gin.H{
		"title":       "我的桌面",
		"ginVersion":  ginVersion,
		"osVersion":   osVersion,
		"os":          os,
		"currentTime": currentTime,
		"user":        user,
	})
}

func LogOut(c *gin.Context) {
	if Services.LogOut(c) {
		c.Header("Cache-Control", "no-cache,no-store")
		c.Redirect(http.StatusMovedPermanently, "/admin")
	}
}

// @Summer 用户列表
func UserList(c *gin.Context) {
	c.HTML(e.SUCCESS, "admin/user.html", gin.H{
		"title": "用户列表",
	})
}

// @Summer 用户列表API
func UserData(c *gin.Context) {
	page := com.StrTo(c.Query("page")).MustInt()
	data := make(map[string]interface{})
	data["list"] = Admin.Users(page)
	data["count"] = e.GetPageNum(Admin.GetUserTotal())
	e.Success(c, "用户列表", data)
}

// @Summer 添加/编辑用户
func AddUser(c *gin.Context) {
	code, msg := Services.AddUser(c)
	e.SendRes(c, code, msg, "")
}

// @Summer 获取单个用户信息
func GetUser(c *gin.Context) {
	_, msg, data := Services.GetUser(c)
	e.Success(c, msg, data)
}

// @Summer 网站信息
func SiteInfo(c *gin.Context) {
	c.HTML(e.SUCCESS, "admin/site.html", gin.H{
		"title": "网站信息",
	})
}

// @Summer 添加/编辑网站信息
func AddSite(c *gin.Context) {
	code, msg := Services.AddSite(c)
	e.SendRes(c, code, msg, "")
}

// @Summer 获取站点信息
func GetSite(c *gin.Context) {
	siteRes := Services.GetSite()
	e.Success(c, "获取站点信息", siteRes)
}

// @Summer 编辑用户信息
func UpdateUser(c *gin.Context) {
	code, msg := Services.EditUser(c)
	e.SendRes(c, code, msg, "")
}
