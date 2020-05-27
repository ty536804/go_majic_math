package Controller

import (
	"elearn100/Model/Article"
	"elearn100/Model/Banner"
	"elearn100/Pkg/e"
	"elearn100/Services"
	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	c.HTML(e.SUCCESS, "index/index.html", gin.H{
		"title": "首页",
	})
}

func About(c *gin.Context) {
	var data = make(map[string]interface{})
	data["banner"] = Banner.GetBannerData(1)
	c.HTML(e.SUCCESS, "index/about.html", gin.H{
		"title": "关于我们",
	})
}

// @Summer 首页
func FrontEnd(c *gin.Context) {
	var data = make(map[string]interface{})
	list := Article.GetArticles(1, data)
	if len(list) > 3 {
		list = list[0:3]
	}

	data["list"] = list
	data["nav"] = Services.GetMenu()
	data["banner"] = Banner.GetBannerData(1)
	data["magic"] = Services.GetSingle(1)
	e.Success(c, "首页", data)
}

// @Summer课程体系
func Subject(c *gin.Context) {
	c.HTML(e.SUCCESS, "index/subject.html", gin.H{
		"title": "课程体系",
	})
}

// @Summer教研教学
func Research(c *gin.Context) {
	c.HTML(e.SUCCESS, "index/research.html", gin.H{
		"title": "教研教学",
	})
}

// @Summer AI学习平台
func Learn(c *gin.Context) {
	c.HTML(e.SUCCESS, "index/learn.html", gin.H{
		"title": "ai学习平台",
	})
}

// @Summer OMO模式
func Omo(c *gin.Context) {
	c.HTML(e.SUCCESS, "index/omo.html", gin.H{
		"title": "OMO模式",
	})
}

// @Summer全国校区
func Campus(c *gin.Context) {
	c.HTML(e.SUCCESS, "index/campus.html", gin.H{
		"title": "全国校区",
	})
}
