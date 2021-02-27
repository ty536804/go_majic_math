package Article

import (
	"elearn100/Model/Article"
	"elearn100/Pkg/e"
	"elearn100/Pkg/setting"
	"elearn100/Services"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

// @Summer文章列表API
func ShowList(c *gin.Context) {
	page := com.StrTo(c.Query("page")).MustInt()
	var data = make(map[string]interface{})
	data["list"] = Article.GetArticles(page, setting.PageSize, data)
	data["count"] = Article.GetArticleTotal()
	data["size"] = setting.PageSize
	e.Success(c, "文章列表", data)
}

// @Summer文章详情Api
func GetArticle(c *gin.Context) {
	c.Request.Body = e.GetBody(c)
	id := com.StrTo(c.PostForm("id")).MustInt()
	var data = make(map[string]interface{})
	data["list"] = Services.RedisGetNavList()
	data["detail"] = Article.GetArticle(id)
	e.Success(c, "文章详情", data)
}

// @Summer文章详情Api
func AddArticle(c *gin.Context) {
	code, msg := Services.AddArticle(c)
	e.SendRes(c, code, msg, "")
}
