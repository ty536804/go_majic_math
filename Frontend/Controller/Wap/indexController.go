package Wap

import (
	"elearn100/Pkg/e"
	"elearn100/Services"
	"github.com/gin-gonic/gin"
	"time"
)

// @Summer 首页
func Index(c *gin.Context) {
	Services.AddVisit(c)
	c.HTML(e.SUCCESS, "wap/index.html", gin.H{
		"title": "首页",
	})
}

// @Summer课程体系
func Subject(c *gin.Context) {
	ver := time.Now().Unix()
	Services.AddVisit(c)
	c.HTML(e.SUCCESS, "wap/subject.html", gin.H{
		"title": "课程体系",
		"time":  ver,
	})
}

// @Summer AI学练系统
func Learn(c *gin.Context) {
	ver := time.Now().Unix()
	Services.AddVisit(c)
	c.HTML(e.SUCCESS, "wap/learn.html", gin.H{
		"title": "AI学联系统",
		"time":  ver,
	})
}

// @Summer omo新模式
func Omo(c *gin.Context) {
	ver := time.Now().Unix()
	Services.AddVisit(c)
	c.HTML(e.SUCCESS, "wap/omo.html", gin.H{
		"title": "omo新模式",
		"time":  ver,
	})
}

// @Summer 加盟授权
func Authorize(c *gin.Context) {
	ver := time.Now().Unix()
	Services.AddVisit(c)
	c.HTML(e.SUCCESS, "wap/join.html", gin.H{
		"title": "加盟授权",
		"time":  ver,
	})
}
