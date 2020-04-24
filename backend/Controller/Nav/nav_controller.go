package Nav

import (
	"elearn100/Services"
	"elearn100/pkg/e"
	"github.com/gin-gonic/gin"
)

// @Summer 导航
func Show(c *gin.Context) {
	c.HTML(e.SUCCESS,"nav/index.html",gin.H{
		"title" : "导航列表",
	})
}

// @Summer 获取一条导航API
func GetNav(c *gin.Context)  {
	data := Services.GetNav(c)
	c.JSON(e.SUCCESS,gin.H{
		"code" : e.SUCCESS,
		"data":data,
	})
}


// @Summer 添加/编辑导航API
func AddNav(c *gin.Context)  {
	code, msg := Services.AddNav(c)
	c.JSON(e.SUCCESS,gin.H{
		"code" : code,
		"msg" : msg,
	})
}

// @Summer 获取多条导航API
func GetNavs(c *gin.Context)  {
	maps := make(map[string]interface{})
	data := Services.GetNavs(maps)
	c.JSON(e.SUCCESS,gin.H{
		"code" : e.SUCCESS,
		"data":data,
	})
}

// @Summer 有效的导航列表
func GetNavList(c *gin.Context)  {
	maps := make(map[string]interface{})
	maps["is_show"] = 1
	data := Services.GetNavs(maps)
	c.JSON(e.SUCCESS,gin.H{
		"code" : e.SUCCESS,
		"data":data,
	})
	e.Success(c,"导航", data)
}