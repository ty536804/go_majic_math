package Admin

import (
	"elearn100/Model/Admin"
	"elearn100/Pkg/e"
	"elearn100/Pkg/setting"
	"elearn100/Services"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"net/http"
)

// @Summer 管理员登录
func Login(c *gin.Context) {
	code, msg := Services.Login(c)
	e.SendRes(c, code, msg, "")
}

func LogOut(c *gin.Context) {
	if Services.LogOut(c) {
		c.Header("Cache-Control", "no-cache,no-store")
		c.Redirect(http.StatusMovedPermanently, "/admin")
	}
}

// @Summer 用户列表API
func UserData(c *gin.Context) {
	page := com.StrTo(c.Query("page")).MustInt()
	data := make(map[string]interface{})
	data["list"] = Admin.Users(page)
	data["count"] = Admin.GetUserTotal()
	data["size"] = setting.PageSize
	e.Success(c, "用户列表", data)
}

// @Summer 添加/编辑用户
func AddUser(c *gin.Context) {
	code, msg := Services.AddUser(c)
	lastId := Services.GetLastUser()
	e.SendRes(c, code, msg, lastId)
}

// @Summer 获取单个用户信息
func GetUser(c *gin.Context) {
	_, msg, data := Services.GetUser(c)
	e.Success(c, msg, data)
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

func DetailsUser(c *gin.Context) {
	isOk, data := Services.DetailsUser(c)
	if isOk != nil {
		e.Error(c, "非法访问", data)
		return
	}
	e.Success(c, "ok", data)
}

// @Summer 编辑短信
func AddSms(c *gin.Context) {
	code, msg := Services.AddSms(c)
	e.SendRes(c, code, msg, "")
}

// @Summer 获取短信配置信息
func GetSms(c *gin.Context) {
	siteRes := Services.GetSms()
	e.Success(c, "获取站点信息", siteRes)
}
