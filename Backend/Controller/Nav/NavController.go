package Nav

import (
	"elearn100/Pkg/e"
	"elearn100/Services"
	"github.com/gin-gonic/gin"
)

// @Summer 导航
func Show(c *gin.Context) {
	c.HTML(e.SUCCESS, "nav/index.html", gin.H{
		"title": "导航列表",
	})
}

// @Summer 获取一条导航API
func GetNav(c *gin.Context) {
	data := Services.GetNav(c)
	e.Success(c, "获取一条导航", data)
}

// @Summer 添加/编辑导航API
func AddNav(c *gin.Context) {
	code, msg := Services.AddNav(c)
	e.SendRes(c, code, msg, "")
}

// @Summer 获取多条导航API
func GetNavs(c *gin.Context) {
	maps := make(map[string]interface{})
	data := Services.GetNavs(maps)
	e.SendRes(c, e.SUCCESS, "获取多条导航", data)
}

// @Summer 有效的导航列表
func GetNavList(c *gin.Context) {
	var data = make(map[string]interface{})
	data["menu"] = Services.GetMenu()
	data["site"] = Services.GetSite()
	data["ercode"] = Services.GetBanner("ercode")
	e.Success(c, "导航", data)
}
