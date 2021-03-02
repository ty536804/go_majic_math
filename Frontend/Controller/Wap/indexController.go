package Wap

import (
	"elearn100/Model/WeChat"
	"elearn100/Pkg/e"
	"elearn100/Pkg/setting"
	"elearn100/Services"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"strconv"
	"strings"
	"time"
)

var baseUrl = "http://www.mofashuxue.com/"

// @Summer 首页
func Index(c *gin.Context) {
	Services.AddVisit(c, baseUrl+"wap")
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
	//强化核心竞争力
	data["indexJzl"] = Services.RedisGetOneBanner(1, 1, "jzl", "wapIndexJzl")
	data["index_cz"] = Services.RedisGetSingleByOne(1, 1, "成长", "subjectCz")
	data["zZBg"] = Services.RedisGetOneBanner(1, 1, "zZBg", "indexZZBg")
	//新闻
	data["index_news"] = Services.RedisGetArticles(1, 3, "indexNew")
	c.HTML(e.SUCCESS, "wap/index.html", gin.H{
		"title": "首页",
		"data":  data,
	})
}

// @Summer课程体系
func Subject(c *gin.Context) {
	//Services.AddVisit(c, baseUrl+"sub")
	data := make(map[string]interface{})
	//六大教学体系
	data["subject_ys"] = Services.RedisGetSingleByOne(3, 1, "数量运算", "subjectYS")
	data["subject_vs"] = Services.RedisGetSingleByOne(3, 1, "序与比较", "subjectVs")
	data["subject_fw"] = Services.RedisGetSingleByOne(3, 1, "空间方位", "subjectFw")
	data["subject_fl"] = Services.RedisGetSingleByOne(3, 1, "形色分类", "subjectFl")
	data["subject_gx"] = Services.RedisGetSingleByOne(3, 1, "对应关系", "subjectGx")
	data["subject_fx"] = Services.RedisGetSingleByOne(3, 1, "逻辑分析", "subjectFx")
	c.HTML(e.SUCCESS, "wap/subject.html", gin.H{
		"title": "课程体系",
		"data":  data,
	})
}

// @Summer AI学练系统
func Learn(c *gin.Context) {
	ver := time.Now().Unix()
	Services.AddVisit(c, baseUrl+"le")
	c.HTML(e.SUCCESS, "wap/learn.html", gin.H{
		"title": "AI学联系统",
		"time":  ver,
	})
}

// @Summer omo新模式
func Omo(c *gin.Context) {
	ver := time.Now().Unix()
	Services.AddVisit(c, baseUrl+"om")
	c.HTML(e.SUCCESS, "wap/omo.html", gin.H{
		"title": "omo新模式",
		"time":  ver,
	})
}

// @Summer 加盟授权
func Authorize(c *gin.Context) {
	ver := time.Now().Unix()
	Services.AddVisit(c, baseUrl+"authorize")
	c.HTML(e.SUCCESS, "wap/join.html", gin.H{
		"title": "加盟授权",
		"time":  ver,
	})
}

// 视频列表
func VideoList(c *gin.Context) {
	page := com.StrTo(c.Query("page")).MustInt()
	data := make(map[string]interface{})
	data["is_show"] = 1
	data["list"] = Services.GetMaterials(page, data)
	data["count"] = e.GetPageNum(Services.GetTotalMaterials())
	data["size"] = setting.PageSize
	c.HTML(e.SUCCESS, "wap/videoList.html", gin.H{
		"title": "视频列表",
		"data":  data,
	})
}

//视频播放
func Video(c *gin.Context) {
	id := com.StrTo(c.Query("id")).MustInt()
	video := Services.GetMaterial(id)
	c.HTML(e.SUCCESS, "wap/video.html", gin.H{
		"title": "视频",
		"video": video,
	})
}

func CheckVideoPwd(c *gin.Context) {
	id := com.StrTo(c.PostForm("id")).MustInt()
	videoPwd := com.StrTo(c.PostForm("video_pwd")).String()
	if id < 1 {
		e.Error(c, "ID不能为空", "")
		return
	}
	video := Services.GetMaterial(id)

	if video.Code != videoPwd {
		e.Error(c, "视频播放码不正确", "")
		return
	}
	data := make(map[string]interface{})
	data["url"] = baseUrl + "videoDetail/id?=" + strconv.Itoa(id)
	uuid := strings.Split(strings.Replace(c.Request.RemoteAddr, ".", "", -1), ":")[0]
	uid, _ := strconv.Atoi(uuid)
	data["user_id"] = uid
	WeChat.AddLook(data)
	e.Success(c, "视频", video)
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
