package Campus

import (
	"elearn100/Pkg/e"
	"elearn100/Services"
	"github.com/gin-gonic/gin"
)

// @Summer 全国校区
func Index(c *gin.Context) {
	c.HTML(e.SUCCESS, "campus/index.html", gin.H{
		"title": "全国校区",
	})
}

// @Summer 获取全国校区API
func GetCampus(c *gin.Context) {
	e.Success(c, "全国校区", Services.GetCampuses(c))
}

// @Summer 获取单个校区API
func DetailCampus(c *gin.Context) {
	e.Success(c, "校区详情", Services.DetailCampus(c))
}

// @Summer 获取全国校区API 带缓冲区的
func GetCampuses(c *gin.Context) {
	e.Success(c, "全国校区", Services.GetCampus(c))
}

// @Summer 获取全国校区API 带缓冲区的
func AddCampuses(c *gin.Context) {
	code, msg := Services.AddCampus(c)
	e.SendRes(c, code, msg, "")
}
