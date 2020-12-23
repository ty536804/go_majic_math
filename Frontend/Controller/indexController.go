package Controller

import (
	"elearn100/Model/Article"
	"elearn100/Model/Banner"
	"elearn100/Model/Single"
	"elearn100/Pkg/e"
	"elearn100/Pkg/setting"
	"elearn100/Services"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

var baseUrl = "http://www.mofashuxue.com/"

// @Summer 首页
func Index(c *gin.Context) {
	Services.AddVisit(c, baseUrl)
	data := make(map[string]interface{})
	LayoutParam(data)
	data["banner"] = Banner.GetBannerData(1, 1, "banner")
	//魔法数学专注于3—12岁少儿数理思维教育
	data["thought"] = Single.GetSingleByOne(1, 1, "thought")
	//什么是数理思维？
	data["why"] = Single.GetSingleByOne(1, 1, "why")
	data["what"] = Banner.GetBannerData(1, 1, "why")
	//为什么要学习数理思维？
	data["index_one"] = Banner.GetOneBanner(1, 1, "01")
	data["index_two"] = Banner.GetOneBanner(1, 1, "02")
	data["index_three"] = Banner.GetOneBanner(1, 1, "03")
	data["siwei"] = Banner.GetOneBanner(1, 1, "数理思维")
	//张梅玲
	data["index_zml"] = Banner.GetOneBanner(1, 1, "张梅玲")
	//魔法数学 颠覆旧数学认知
	data["index_df"] = Banner.GetOneBanner(1, 1, "颠覆BK")
	data["index_vs"] = Banner.GetOneBanner(1, 1, "vs")
	data["index_con"] = Single.GetSingleByOne(1, 1, "vsCon")
	//全新OMO教育模式 形成全场景闭环教学空间
	data["index_omol"] = Banner.GetOneBanner(1, 1, "omo_l")
	data["index_omor"] = Banner.GetOneBanner(1, 1, "omo_r")
	//data["index_omom"] = indexOmoM

	data["index_cz"] = Single.GetSingleByOne(1, 1, "成长")
	data["index_power_l"] = Banner.GetOneBanner(1, 1, "power_con_l")
	data["index_power_m"] = Single.GetSingleByOne(1, 1, "power_con_m")
	data["index_power_r"] = Banner.GetOneBanner(1, 1, "power_con_r")
	//魔法数学课程体系 为少儿提供高品质的学习方案
	data["index_fanan"] = Single.GetSingleByOne(1, 1, "方案")
	data["planning"] = Single.GetSingleByOne(1, 1, "课程规划")
	data["course"] = Banner.GetBannerData(1, 1, "课时规划")
	//趣味课堂 让孩子重拾数学乐趣
	data["index_happy"] = Single.GetAllSingle(1, 1, "happy")
	//新闻
	data["index_news"] = Article.GetArticles(1, 3, make(map[string]interface{}))
	data["newBg"] = Banner.GetOneBanner(1, 1, "newBg")
	//强化核心竞争力
	data["indexJzl"] = Banner.GetOneBanner(1, 1, "jzl")
	data["zZBg"] = Banner.GetOneBanner(1, 1, "zZBg")

	c.HTML(e.SUCCESS, "index/index.html", gin.H{
		"title": "首页",
		"data":  data,
	})
}

// @Summer 关于魔数
func About(c *gin.Context) {
	_url := baseUrl + "about"
	Services.AddVisit(c, _url)
	data := make(map[string]interface{})
	LayoutParam(data)
	data["banner"] = Banner.GetBannerData(2, 1, "banner")
	data["aboutBk"] = Banner.GetOneBanner(2, 1, "aboutBk")
	data["swg"] = Single.GetSingleByOne(2, 1, "swg")
	data["child"] = Banner.GetOneBanner(2, 1, "卓越的孩子")
	data["midBk"] = Banner.GetOneBanner(2, 1, "midBk")
	data["midSlogan"] = Banner.GetOneBanner(2, 1, "midSlogan")
	data["brandBk"] = Banner.GetOneBanner(2, 1, "brandBk")
	data["brand"] = Banner.GetBannerData(2, 1, "brand")
	data["human"] = Banner.GetBannerData(2, 1, "human")
	data["brandBg"] = Banner.GetOneBanner(2, 1, "brandBg")
	data["year"] = Banner.GetBannerData(2, 1, "year")
	//Services.AddVisit(c, baseUrl+"about")
	c.HTML(e.SUCCESS, "index/about.html", gin.H{
		"title": "关于魔数",
		"data":  data,
	})
}

// @Summer课程体系
func Subject(c *gin.Context) {
	_url := baseUrl + "subject"
	Services.AddVisit(c, _url)
	data := make(map[string]interface{})
	LayoutParam(data)
	data["banner"] = Banner.GetBannerData(3, 1, "banner")
	//强化核心竞争力
	data["indexJzl"] = Banner.GetOneBanner(1, 1, "jzl")
	data["index_cz"] = Single.GetSingleByOne(1, 1, "成长")
	//魔法数学课程体系 为少儿提供高品质的学习方案
	data["index_fanan"] = Single.GetSingleByOne(1, 1, "方案")
	data["planning"] = Single.GetSingleByOne(1, 1, "课程规划")
	data["course"] = Banner.GetBannerData(1, 1, "课时规划")
	//趣味课堂 让孩子重拾数学乐趣
	data["index_happy"] = Single.GetAllSingle(1, 1, "happy")
	data["happy_warp"] = Banner.GetOneBanner(2, 1, "happy_warp")
	//六大教学体系 促进思维发展
	data["cjrBk"] = Banner.GetOneBanner(2, 1, "cjrBk")
	data["cjcBk"] = Banner.GetOneBanner(2, 1, "cjcBk")
	data["cjlBk"] = Banner.GetOneBanner(2, 1, "cjlBk")
	//六大教学体系
	data["subject_ys"] = Single.GetSingleByOne(3, 1, "数量运算")
	data["subject_vs"] = Single.GetSingleByOne(3, 1, "序与比较")
	data["subject_fw"] = Single.GetSingleByOne(3, 1, "空间方位")
	data["subject_fl"] = Single.GetSingleByOne(3, 1, "形色分类")
	data["subject_gx"] = Single.GetSingleByOne(3, 1, "对应关系")
	data["subject_fx"] = Single.GetSingleByOne(3, 1, "逻辑分析")
	c.HTML(e.SUCCESS, "index/subject.html", gin.H{
		"title": "课程体系",
		"data":  data,
	})
}

// @Summer教研教学
func Research(c *gin.Context) {
	_url := baseUrl + "research"
	Services.AddVisit(c, _url)
	data := make(map[string]interface{})
	LayoutParam(data)
	data["banner"] = Banner.GetBannerData(4, 1, "banner")
	data["celebrity"] = Single.GetAllSingle(4, 1, "celebrity")
	//魔法数学教材 让独立思考成为孩子的习惯
	data["xiguai"] = Single.GetSingleByOne(4, 1, "xiguai")
	data["course"] = Banner.GetBannerData(4, 1, "course")
	//多种教具配合教材使用 让学习效果事半功倍
	data["double"] = Single.GetSingleByOne(4, 1, "double")
	//“魔法杯”思维盛典 让思维力的佼佼者脱颖而出
	data["mfbCon"] = Single.GetSingleByOne(4, 1, "mfb")
	data["mfb"] = Banner.GetBannerData(4, 1, "mfb")
	data["big"] = Banner.GetOneBanner(4, 1, "大合影")
	data["show"] = Banner.GetOneBanner(4, 1, "A86")

	c.HTML(e.SUCCESS, "index/research.html", gin.H{
		"title": "教研教学",
		"data":  data,
	})
}

// @Summer AI学习平台
func Learn(c *gin.Context) {
	_url := baseUrl + "learn"
	Services.AddVisit(c, _url)
	data := make(map[string]interface{})
	LayoutParam(data)
	data["banner"] = Banner.GetBannerData(5, 1, "banner")
	//全域成长 智慧学习
	data["learn"] = Single.GetSingleByOne(5, 1, "智慧学习")
	data["online"] = Single.GetAllSingle(5, 1, "onlinesys")

	c.HTML(e.SUCCESS, "index/learn.html", gin.H{
		"title": "智学系统",
		"data":  data,
	})
}

// @Summer 智学系统
func Omo(c *gin.Context) {
	_url := baseUrl + "omo"
	Services.AddVisit(c, _url)
	data := make(map[string]interface{})
	LayoutParam(data)
	data["banner"] = Banner.GetBannerData(6, 1, "banner")
	//全新OMO教育模式 形成全场景闭环教学空间
	data["index_omol"] = Banner.GetOneBanner(1, 1, "omo_l")
	data["index_omor"] = Banner.GetOneBanner(1, 1, "omo_r")
	//OMO推动新一代教育变革 为教育赋能
	data["energize"] = Single.GetSingleByOne(6, 1, "energize")
	data["fn"] = Single.GetAllSingle(6, 1, "fn")
	c.HTML(e.SUCCESS, "index/omo.html", gin.H{
		"title": "OMO模式",
		"data":  data,
	})
}

// @Summer全国校区
func Campus(c *gin.Context) {
	_url := baseUrl + "campus"
	Services.AddVisit(c, _url)
	data := make(map[string]interface{})
	LayoutParam(data)
	data["school_bg"] = Banner.GetOneBanner(7, 1, "school_bg")
	c.HTML(e.SUCCESS, "index/campus.html", gin.H{
		"title": "全国校区",
		"data":  data,
	})
}

// @Summer 新闻动态
func News(c *gin.Context) {
	_url := baseUrl + "news"
	Services.AddVisit(c, _url)
	data := make(map[string]interface{})
	LayoutParam(data)
	data["banner"] = Banner.GetBannerData(11, 1, "banner")
	where := make(map[string]interface{})
	where["is_show"] = 1
	data["list"] = Article.GetArticles(0, setting.PageSize, where)
	data["count"] = e.GetPageNum(Article.GetArticleTotal())
	c.HTML(e.SUCCESS, "index/new.html", gin.H{
		"title": "新闻动态",
		"data":  data,
	})
}

// @Summer 新闻动态列表
func NewList(c *gin.Context) {
	page := com.StrTo(c.Query("page")).MustInt()
	var data = make(map[string]interface{})
	data["is_show"] = 1
	data["list"] = Article.GetArticles(page, setting.PageSize, data)
	data["count"] = e.GetPageNum(Article.GetArticleTotal())
	e.Success(c, "首页", data)
}

// @Summer 新闻详情
func NewDetail(c *gin.Context) {
	data := make(map[string]interface{})
	LayoutParam(data)
	id := com.StrTo(c.DefaultQuery("id", "0")).MustInt()
	_url := baseUrl + "detail/?id=" + string(id)
	Services.AddVisit(c, _url)
	c.HTML(e.SUCCESS, "index/detail.html", gin.H{
		"title":  "新闻详情",
		"detail": Article.GetArticle(id),
		"data":   data,
	})
}

// @Summer 加盟授权
func Authorize(c *gin.Context) {
	_url := baseUrl + "join"
	Services.AddVisit(c, _url)
	data := make(map[string]interface{})
	LayoutParam(data)
	data["banner"] = Banner.GetBannerData(8, 1, "banner")
	c.HTML(e.SUCCESS, "index/join.html", gin.H{
		"title": "加盟授权",
		"data":  data,
	})
}

// @Summer 加盟授权
func Down(c *gin.Context) {
	_url := baseUrl + "down"
	Services.AddVisit(c, _url)
	data := make(map[string]interface{})
	LayoutParam(data)
	data["banner"] = Banner.GetBannerData(1, 1, "banner")
	//魔数小奇兵APP
	data["energize"] = Single.GetSingleByOne(9, 1, "魔数小奇兵APP")
	data["flower"] = Banner.GetBannerData(9, 1, "flower")
	data["wisdom"] = Banner.GetBannerData(9, 1, "智慧")
	c.HTML(e.SUCCESS, "index/down.html", gin.H{
		"title": "下载中心",
		"data":  data,
	})
}

func GetWeChat(c *gin.Context) {
	start := com.StrTo(c.DefaultQuery("start", "0")).MustInt()
	end := com.StrTo(c.DefaultQuery("end", "0")).MustInt()
	if start >= 0 && end > 0 {
		Services.GetArticle(start, end)
	}
}

// 所有页面用到的内容
func LayoutParam(data map[string]interface{}) map[string]interface{} {
	headslogan := Banner.GetOneBanner(1, 1, "headslogan")
	headLogo := Banner.GetOneBanner(1, 1, "headLogo")
	mfsx := Banner.GetOneBanner(1, 1, "mfsx")
	msxqb := Banner.GetOneBanner(1, 1, "msxqb")
	data["menu"] = Services.GetMenu()
	data["siteInfo"] = Services.GetSite()
	data["mfsx"] = mfsx
	data["msxqb"] = msxqb
	data["headslogan"] = headslogan
	data["headLogo"] = headLogo
	return data
}
