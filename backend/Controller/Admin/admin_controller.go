package Admin

import (
	"elearn100/Model/Admin"
	"elearn100/Services"
	"elearn100/pkg/e"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"net/http"
	"runtime"
	"time"
)

// @Summer 管理员登录
func Login(c *gin.Context) {
	var data interface{}
	code, err := Services.Login(c)
	if code == e.ERROR {
		e.Error(c,err,data)
	} else {
		e.Success(c,"登录成功",err)
	}
}

// @Summer 后端首页
func Show(c *gin.Context)  {
	c.HTML(e.SUCCESS,"admin/home.html",gin.H{
		"title":"易学教育",
	})
}

// @Summer 后端首页详情内容
func Welcome(c *gin.Context)  {
	ginVersion := gin.Version
	osVersion := runtime.Version()
	os := runtime.GOOS
	currentTime := time.Now().Format("2006:01:02 15:04:05")

	c.HTML(e.SUCCESS,"admin/welcome.html",gin.H{
		"title":"content",
		"ginVersion":ginVersion,
		"osVersion":osVersion,
		"os":os,
		"currentTime":currentTime,
	})
}

func LogOut(c *gin.Context)  {
	if isOk, _:= e.GetVal("token");isOk {
		e.DelVal("token")
	}
	c.Header("Cache-Control","no-cache,no-store")
	c.Redirect(http.StatusMovedPermanently, "/admin")
}

// @Summer 用户列表
func UserList(c *gin.Context)  {
	c.HTML(e.SUCCESS,"admin/user.html",gin.H{
		"title":"用户列表",
	})
}

// @Summer 用户列表API
func UserData(c *gin.Context)  {
	page := com.StrTo(c.Query("page")).MustInt()
	data := make(map[string]interface{})
	data["list"] = Admin.Users(page)
	data["count"] =  e.GetPageNum(Admin.GetUserTotal())
	c.JSON(e.SUCCESS,gin.H{
		"code" : e.SUCCESS,
		"msg" : "用户列表",
		"data" : data,
	})
}

// @Summer 添加/编辑用户
func AddUser(c *gin.Context)  {
	code, err := Services.AddUser(c)
	data := make(map[string]interface{})
	if code == e.SUCCESS {
		data["count"] = e.GetPageNum(Admin.GetUserTotal())
	}
	c.JSON(e.SUCCESS,gin.H{
		"code" : code,
		"msg" : err,
		"data" : data,
	})
}

// @Summer 获取单个用户信息
func GetUser(c *gin.Context) {
	code, err,data := Services.GetUser(c)
	c.JSON(e.SUCCESS,gin.H{
		"code" : code,
		"msg" : err,
		"data" : data,
	})
}

// @Summer 网站信息
func SiteInfo(c *gin.Context)  {
	c.HTML(e.SUCCESS,"admin/site.html",gin.H{
		"title":"网站信息",
	})
}

// @Summer 添加/编辑网站信息
func AddSite(c *gin.Context)  {
	code, err := Services.AddSite(c)
	c.JSON(e.SUCCESS,gin.H{
		"code" : code,
		"msg" : err,
	})
}

// @Summer 获取站点信息
func GetSite(c *gin.Context)  {
	siteRes := Services.GetSite()
	c.JSON(e.SUCCESS,gin.H{
		"code" : e.SUCCESS,
		"msg" : "获取站点信息",
		"data" : siteRes,
	})
}