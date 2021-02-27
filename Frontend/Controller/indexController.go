package Controller

import (
	"elearn100/Model/Article"
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
	data["banner"] = Services.RedisGetBannerList(1, 1, "banner", "indexBanner")
	//魔法数学专注于3—12岁少儿数理思维教育
	data["thought"] = Services.RedisGetSingleByOne(1, 1, "thought", "indexThought")
	//什么是数理思维？
	data["why"] = Services.RedisGetSingleByOne(1, 1, "why", "indexWhy")
	data["what"] = Services.RedisGetBannerList(1, 1, "why", "indexWhat")
	//为什么要学习数理思维？
	data["index_one"] = Services.RedisGetOneBanner(1, 1, "01", "indexOne")
	data["index_two"] = Services.RedisGetOneBanner(1, 1, "02", "indexTwo")
	data["index_three"] = Services.RedisGetOneBanner(1, 1, "03", "indexThree")
	data["siwei"] = Services.RedisGetOneBanner(1, 1, "数理思维", "indexSiWei")
	//张梅玲
	data["index_zml"] = Services.RedisGetOneBanner(1, 1, "张梅玲", "indexZml")
	//魔法数学 颠覆旧数学认知
	data["index_df"] = Services.RedisGetOneBanner(1, 1, "颠覆BK", "indexDf")
	data["index_vs"] = Services.RedisGetOneBanner(1, 1, "vs", "indexVs")
	data["index_con"] = Services.RedisGetSingleByOne(1, 1, "vsCon", "indexVsCon")
	//全新OMO教育模式 形成全场景闭环教学空间
	data["index_omol"] = Services.RedisGetOneBanner(1, 1, "omo_l", "indexOmoL")
	data["index_omor"] = Services.RedisGetOneBanner(1, 1, "omo_r", "indexOmoR")
	//data["index_omom"] = indexOmoM

	data["index_cz"] = Services.RedisGetSingleByOne(1, 1, "成长", "indexCz")
	//data["index_power_l"] = Services.RedisGetOneBanner(1, 1, "power_con_l","indexPowerL")
	data["index_power_m"] = Services.RedisGetSingleByOne(1, 1, "power_con_m", "indexPowerM")
	//data["index_power_r"] = Services.RedisGetOneBanner(1, 1, "power_con_r","indexPowerR")
	//魔法数学课程体系 为少儿提供高品质的学习方案
	data["index_fanan"] = Services.RedisGetSingleByOne(1, 1, "方案", "indexFanAn")
	data["planning"] = Services.RedisGetSingleByOne(1, 1, "课程规划", "indexPlanning")
	data["course"] = Services.RedisGetBannerList(1, 1, "课时规划", "indexCourse")
	//趣味课堂 让孩子重拾数学乐趣
	data["index_happy"] = Services.RedisGetAllSingle(1, 1, "happy", "indexHappy")
	//新闻
	data["index_news"] = Services.RedisGetArticles(1, 3, "indexNew")
	data["newBg"] = Services.RedisGetOneBanner(1, 1, "newBg", "indexNewBg")
	//强化核心竞争力
	data["indexJzl"] = Services.RedisGetOneBanner(1, 1, "jzl", "indexJzl")
	data["zZBg"] = Services.RedisGetOneBanner(1, 1, "zZBg", "indexZZBg")

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
	data["banner"] = Services.RedisGetBannerList(2, 1, "banner", "aboutBanner")
	data["aboutBk"] = Services.RedisGetOneBanner(2, 1, "aboutBk", "aboutBk")
	data["swg"] = Services.RedisGetSingleByOne(2, 1, "swg", "aboutSwg")
	data["child"] = Services.RedisGetOneBanner(2, 1, "卓越的孩子", "aboutChild")
	data["midBk"] = Services.RedisGetOneBanner(2, 1, "midBk", "aboutMidBk")
	data["midSlogan"] = Services.RedisGetOneBanner(2, 1, "midSlogan", "aboutMidSlogan")
	data["brandBk"] = Services.RedisGetOneBanner(2, 1, "brandBk", "aboutBrandBk")
	data["brand"] = Services.RedisGetBannerList(2, 1, "brand", "aboutBrand")
	data["human"] = Services.RedisGetBannerList(2, 1, "human", "aboutHuman")
	data["brandBg"] = Services.RedisGetOneBanner(2, 1, "brandBg", "aboutBrandBg")
	data["year"] = Services.RedisGetBannerList(2, 1, "year", "aboutYear")
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
	data["banner"] = Services.RedisGetBannerList(3, 1, "banner", "subjectBanner")
	//强化核心竞争力
	data["indexJzl"] = Services.RedisGetOneBanner(1, 1, "jzl", "subjectJzl")
	data["index_cz"] = Services.RedisGetSingleByOne(1, 1, "成长", "subjectCz")
	//魔法数学课程体系 为少儿提供高品质的学习方案
	data["index_fanan"] = Services.RedisGetSingleByOne(1, 1, "方案", "subjectFanAn")
	data["planning"] = Services.RedisGetSingleByOne(1, 1, "课程规划", "subjectPlaning")
	data["course"] = Services.RedisGetBannerList(1, 1, "课时规划", "subjectCourse")
	//趣味课堂 让孩子重拾数学乐趣
	data["index_happy"] = Services.RedisGetAllSingle(1, 1, "happy", "subjectHappy")
	data["happy_warp"] = Services.RedisGetOneBanner(2, 1, "happy_warp", "subjectHappy")
	//六大教学体系 促进思维发展
	data["cjrBk"] = Services.RedisGetOneBanner(2, 1, "cjrBk", "subjectCjrBk")
	data["cjcBk"] = Services.RedisGetOneBanner(2, 1, "cjcBk", "subjectCjcBk")
	data["cjlBk"] = Services.RedisGetOneBanner(2, 1, "cjlBk", "subjectCjlBk")
	//六大教学体系
	data["subject_ys"] = Services.RedisGetSingleByOne(3, 1, "数量运算", "subjectYS")
	data["subject_vs"] = Services.RedisGetSingleByOne(3, 1, "序与比较", "subjectVs")
	data["subject_fw"] = Services.RedisGetSingleByOne(3, 1, "空间方位", "subjectFw")
	data["subject_fl"] = Services.RedisGetSingleByOne(3, 1, "形色分类", "subjectFl")
	data["subject_gx"] = Services.RedisGetSingleByOne(3, 1, "对应关系", "subjectGx")
	data["subject_fx"] = Services.RedisGetSingleByOne(3, 1, "逻辑分析", "subjectFx")
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
	data["banner"] = Services.RedisGetBannerList(4, 1, "banner", "researchBanner")
	data["celebrity"] = Services.RedisGetAllSingle(4, 1, "celebrity", "researchCelebrity")
	//魔法数学教材 让独立思考成为孩子的习惯
	data["xiguai"] = Services.RedisGetSingleByOne(4, 1, "xiguai", "researchXiGuai")
	data["course"] = Services.RedisGetBannerList(4, 1, "course", "researchCourse")
	//多种教具配合教材使用 让学习效果事半功倍
	data["double"] = Services.RedisGetSingleByOne(4, 1, "double", "researchDouble")
	//“魔法杯”思维盛典 让思维力的佼佼者脱颖而出
	data["mfbCon"] = Services.RedisGetSingleByOne(4, 1, "mfb", "researchMfb")
	data["mfb"] = Services.RedisGetBannerList(4, 1, "mfb", "researchMfb")
	data["big"] = Services.RedisGetOneBanner(4, 1, "大合影", "researchBig")
	data["show"] = Services.RedisGetOneBanner(4, 1, "A86", "researchShow")

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
	data["banner"] = Services.RedisGetBannerList(5, 1, "banner", "learnBanner")
	//全域成长 智慧学习
	data["learn"] = Services.RedisGetSingleByOne(5, 1, "智慧学习", "learnLearn")
	data["online"] = Services.RedisGetAllSingle(5, 1, "onlinesys", "learnOnlineSys")

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
	data["banner"] = Services.RedisGetBannerList(6, 1, "banner", "omoBanner")
	//全新OMO教育模式 形成全场景闭环教学空间
	data["index_omol"] = Services.RedisGetOneBanner(1, 1, "omo_l", "omoL")
	data["index_omor"] = Services.RedisGetOneBanner(1, 1, "omo_r", "omoR")
	//OMO推动新一代教育变革 为教育赋能
	data["energize"] = Services.RedisGetSingleByOne(6, 1, "energize", "omoEnergize")
	data["fn"] = Services.RedisGetAllSingle(6, 1, "fn", "omoFn")
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
	data["school_bg"] = Services.RedisGetOneBanner(7, 1, "school_bg", "campusSchoolBg")
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
	data["banner"] = Services.RedisGetBannerList(11, 1, "banner", "newBanner")
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
	data["banner"] = Services.RedisGetBannerList(8, 1, "banner", "authorizeBanner")
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
	data["banner"] = Services.RedisGetBannerList(9, 1, "banner", "downBanner")
	//魔数小奇兵APP
	data["energize"] = Services.RedisGetSingleByOne(9, 1, "魔数小奇兵APP", "downEnergize")
	data["flower"] = Services.RedisGetBannerList(9, 1, "flower", "downFlower")
	data["wisdom"] = Services.RedisGetBannerList(9, 1, "智慧", "downWisdom")
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
	headslogan := Services.RedisGetOneBanner(1, 1, "headslogan", "headSlogan")
	headLogo := Services.RedisGetOneBanner(1, 1, "headLogo", "headLogo")
	mfsx := Services.RedisGetOneBanner(1, 1, "mfsx", "mfsx")
	msxqb := Services.RedisGetOneBanner(1, 1, "msxqb", "msxqb")
	data["menu"] = Services.GetMenu()
	data["siteInfo"] = Services.GetSite()
	data["mfsx"] = mfsx
	data["msxqb"] = msxqb
	data["headslogan"] = headslogan
	data["headLogo"] = headLogo
	return data
}
