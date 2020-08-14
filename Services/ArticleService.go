package Services

import (
	"elearn100/Model/Article"
	"elearn100/Pkg/e"
	"fmt"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

// @Summer 添加文章
func AddArticle(c *gin.Context) (code int, msg string) {
	var data = make(map[string]interface{})
	if err := c.Bind(&c.Request.Body); err != nil {
		fmt.Println(err)
		return e.ERROR, "操作失败"
	}
	id := com.StrTo(c.PostForm("id")).MustInt()
	title := com.StrTo(c.PostForm("title")).String()
	summary := com.StrTo(c.PostForm("summary")).String()
	admin := com.StrTo(c.PostForm("admin")).String()
	content := com.StrTo(c.PostForm("content")).String()
	isShow := com.StrTo(c.PostForm("is_show")).MustInt()
	sort := com.StrTo(c.PostForm("sort")).MustInt()
	hot := com.StrTo(c.PostForm("hot")).MustInt()
	thumbImg := com.StrTo(c.PostForm("thumb_img")).String()
	articleCom := com.StrTo(c.PostForm("com")).String()
	navId := com.StrTo(c.PostForm("nav_id")).MustInt()

	valid := validation.Validation{}
	valid.Required(title, "title").Message("标题不能为空")
	valid.Required(summary, "summary").Message("摘要不能为空")
	valid.Required(isShow, "is_show").Message("选择是否展示")
	valid.Required(content, "content").Message("内容不能为空")
	valid.Required(hot, "hot").Message("选择是否热点")
	valid.Required(admin, "admin").Message("发布者不能为空")
	valid.Required(navId, "nav_id").Message("栏目不能为空")
	//valid.Required(com,"com").Message("来源不能为空")

	if !valid.HasErrors() {
		data["title"] = title
		data["summary"] = summary
		data["is_show"] = isShow
		data["content"] = content
		data["hot"] = hot
		data["sort"] = sort
		data["thumb_img"] = thumbImg
		data["admin"] = admin
		data["com"] = articleCom
		data["nav_id"] = navId

		isOk := false
		if id < 1 {
			isOk = Article.AddArticle(data)
		} else {
			isOk = Article.EditArticle(id, data)
		}
		if isOk {
			return e.SUCCESS, "操作成功"
		}
		return e.ERROR, "操作失败"
	}
	return ViewErr(valid)
}
