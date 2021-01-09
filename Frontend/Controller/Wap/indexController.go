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
	c.HTML(e.SUCCESS, "wap/index.html", gin.H{
		"title": "首页",
	})
}

// @Summer课程体系
func Subject(c *gin.Context) {
	ver := time.Now().Unix()
	Services.AddVisit(c, baseUrl+"sub")
	c.HTML(e.SUCCESS, "wap/subject.html", gin.H{
		"title": "课程体系",
		"time":  ver,
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
