package Article

import (
	"elearn100/Model/Article"
	"elearn100/Services"
	"elearn100/pkg/e"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

// @Summer文章列表页面
func Show(c *gin.Context) {
	c.HTML(e.SUCCESS, "article/index.html", gin.H{
		"title": "文章列表",
	})
}

// @Summer文章列表API
func ShowList(c *gin.Context) {
	page := com.StrTo(c.Query("page")).MustInt()
	var data = make(map[string]interface{})
	data["list"] = Article.GetArticles(page, data)
	data["count"] = e.GetPageNum(Article.GetArticleTotal())
	e.Success(c, "文章列表", data)
}

// @Summer文章详情
func Detail(c *gin.Context) {
	id := com.StrTo(c.DefaultQuery("id", "0")).MustInt()
	c.HTML(e.SUCCESS, "article/detail.html", gin.H{
		"title": "文章详情页面",
		"id":    id,
	})
}

// @Summer文章详情Api
func GetArticle(c *gin.Context) {
	c.Request.Body = e.GetBody(c)
	id := com.StrTo(c.PostForm("id")).MustInt()
	fmt.Print(id)
	data := Article.GetArticle(id)
	e.Success(c, "文章详情", data)
}

// @Summer文章详情Api
func AddArticle(c *gin.Context) {
	code, msg := Services.AddArticle(c)
	e.SendRes(c, code, msg, "")
}
