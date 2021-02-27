package Single

import (
	"elearn100/Model/Single"
	"elearn100/Pkg/e"
	"elearn100/Pkg/setting"
	"elearn100/Services"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

// @Summer 单页列表
func ListData(c *gin.Context) {
	page := com.StrTo(c.Query("page")).MustInt()
	data := make(map[string]interface{})
	data["list"] = Single.GetSingles(page, data)
	data["count"] = Single.GetSingleTotal()
	data["size"] = setting.PageSize
	e.Success(c, "单页列表", data)
}

// @Summer 添加单页
func AddSingle(c *gin.Context) {
	code, msg := Services.AddSingle(c)
	e.Success(c, msg, code)
}

// @Summer 单页详情Api
func GetSingle(c *gin.Context) {
	c.Request.Body = e.GetBody(c)
	id := com.StrTo(c.PostForm("id")).MustInt()
	var data = make(map[string]interface{})
	data["list"] = Services.RedisGetNavList()
	data["detail"] = Single.GetSingle(id)
	e.Success(c, "单页文章详情", data)
}
