package Single

import (
	"elearn100/Model/Single"
	"elearn100/Pkg/e"
	"elearn100/Services"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

func List(c *gin.Context) {
	c.HTML(e.SUCCESS, "single/index.html", gin.H{
		"title": "单页列表",
	})
}

// @Summer 单页列表
func ListData(c *gin.Context) {
	page := com.StrTo(c.PostForm("page")).MustInt()
	data := make(map[string]interface{})
	data["list"] = Single.GetSingles(page, data)
	data["count"] = e.GetPageNum(Single.GetSingleTotal())
	e.Success(c, "单页列表", data)
}

// @Summer 添加单页
func AddSingle(c *gin.Context) {
	code, msg := Services.AddSingle(c)
	e.Success(c, msg, code)
}

// @Summer文章详情Api
func GetSingle(c *gin.Context) {
	c.Request.Body = e.GetBody(c)
	id := com.StrTo(c.PostForm("id")).MustInt()
	var data = make(map[string]interface{})
	data["list"] = Services.GetNavs(data)
	data["detail"] = Single.GetSingle(id)
	e.Success(c, "单页文章详情", data)
}

// @Summer文章详情
func DetailSingle(c *gin.Context) {
	id := com.StrTo(c.DefaultQuery("id", "0")).MustInt()
	c.HTML(e.SUCCESS, "single/detail.html", gin.H{
		"title": "单页详情",
		"id":    id,
	})
}
