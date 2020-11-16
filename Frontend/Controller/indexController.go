package Controller

import (
	"elearn100/Model/Article"
	"elearn100/Model/Banner"
	"elearn100/Pkg/e"
	"elearn100/Services"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

var baseUrl = "http://www.mofashuxue.com/"

func Index(c *gin.Context) {
	Services.AddVisit(c, baseUrl)
	banner := Banner.GetOneBanner(1)
	c.HTML(e.SUCCESS, "index/index.html", gin.H{
		"title":  "首页",
		"banner": banner.Imgurl,
	})
}

// @Summer 首页
func FrontEnd(c *gin.Context) {
	var data = make(map[string]interface{})
	list := Article.GetArticles(1, data)
	if len(list) > 5 {
		list = list[0:4]
	}

	banner := Banner.GetBannerData(1)
	data["list"] = list
	data["banner"] = banner
	data["magic"] = Services.GetCon(1)
	e.Success(c, "首页", data)
}

func About(c *gin.Context) {
	Services.AddVisit(c, baseUrl+"about")
	var data = make(map[string]interface{})
	data["banner"] = Banner.GetBannerData(1)
	c.HTML(e.SUCCESS, "index/about.html", gin.H{
		"title": "关于我们",
	})
}

func AboutData(c *gin.Context) {
	var data = make(map[string]interface{})
	list := Article.GetArticles(1, data)
	if len(list) > 3 {
		list = list[0:3]
	}
	banner := Banner.GetBannerData(2)
	data["banner"] = banner
	data["magic"] = Services.GetCon(2)
	e.Success(c, "关于我们", data)
}

// @Summer课程体系
func Subject(c *gin.Context) {
	Services.AddVisit(c, baseUrl+"subject")
	c.HTML(e.SUCCESS, "index/subject.html", gin.H{
		"title": "课程体系",
	})
}

// @Summer教研教学
func Research(c *gin.Context) {
	Services.AddVisit(c, baseUrl+"research")
	c.HTML(e.SUCCESS, "index/research.html", gin.H{
		"title": "教研教学",
	})
}

// @Summer AI学习平台
func Learn(c *gin.Context) {
	Services.AddVisit(c, baseUrl+"learn")
	c.HTML(e.SUCCESS, "index/ai.html", gin.H{
		"title": "ai学习平台",
	})
}

// @Summer OMO模式
func Omo(c *gin.Context) {
	Services.AddVisit(c, baseUrl+"omo")
	c.HTML(e.SUCCESS, "index/omo.html", gin.H{
		"title": "OMO模式",
	})
}

// @Summer全国校区
func Campus(c *gin.Context) {
	Services.AddVisit(c, baseUrl+"campus")
	c.HTML(e.SUCCESS, "index/campus.html", gin.H{
		"title": "全国校区",
	})
}

// @Summer 新闻动态
func News(c *gin.Context) {
	Services.AddVisit(c, baseUrl+"news")
	c.HTML(e.SUCCESS, "index/new.html", gin.H{
		"title": "新闻动态",
	})
}

// @Summer 新闻动态列表
func NewList(c *gin.Context) {
	page := com.StrTo(c.Query("page")).MustInt()
	var data = make(map[string]interface{})
	data["is_show"] = 1
	data["list"] = Article.GetArticles(page, data)
	data["count"] = e.GetPageNum(Article.GetArticleTotal())
	e.Success(c, "首页", data)
}

// @Summer 新闻详情
func NewDetail(c *gin.Context) {
	id := com.StrTo(c.DefaultQuery("id", "0")).MustInt()
	_url := baseUrl + "detail/?id=" + string(id)
	Services.AddVisit(c, _url)
	c.HTML(e.SUCCESS, "index/detail.html", gin.H{
		"title":  "新闻详情",
		"detail": Article.GetArticle(id),
	})
}

// @Summer 加盟授权
func Authorize(c *gin.Context) {
	Services.AddVisit(c, baseUrl+"join")
	c.HTML(e.SUCCESS, "index/join.html", gin.H{
		"title": "加盟授权",
	})
}

// @Summer 加盟授权数据接口
func JoinData(c *gin.Context) {
	c.Request.Body = e.GetBody(c)
	var data = make(map[string]interface{})
	data["banner"] = Banner.GetData(8, 0)
	data["app"] = Banner.GetData(8, 1)
	data["learn"] = Banner.GetData(8, 2)
	data["mid"] = Banner.GetData(8, 3)
	e.Success(c, "success", data)
}

// @Summer 加盟授权
func Down(c *gin.Context) {
	Services.AddVisit(c, baseUrl+"down")
	c.HTML(e.SUCCESS, "index/down.html", gin.H{
		"title": "APP下载",
	})
}
func GetWeChat(c *gin.Context) {
	Services.GetArticle(0, 5)
}
