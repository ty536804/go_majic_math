package Services

import (
	"elearn100/Model/Single"
	"elearn100/Pkg/e"
	"fmt"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

// @Summer 添加文章
func AddSingle(c *gin.Context) (code int, msg string) {
	var data = make(map[string]interface{})
	if err := c.Bind(&c.Request.Body); err != nil {
		fmt.Println(err)
		return e.ERROR, "操作失败"
	}
	id := com.StrTo(c.PostForm("id")).MustInt()
	name := com.StrTo(c.PostForm("name")).String()
	navId := com.StrTo(c.PostForm("nav_id")).MustInt()
	content := com.StrTo(c.PostForm("content")).String()
	thumbImg := com.StrTo(c.PostForm("thumb_img")).String()
	summary := com.StrTo(c.PostForm("summary")).String()
	tag := com.StrTo(c.PostForm("tag")).String()
	ClientType := com.StrTo(c.PostForm("client_type")).MustInt()

	valid := validation.Validation{}
	valid.Required(name, "name").Message("标题不能为空")
	valid.Required(content, "content").Message("内容不能为空")
	valid.Required(navId, "nav_id").Message("栏目不能为空")
	valid.Required(tag, "tag").Message("标签不能为空")

	if !valid.HasErrors() {
		data["name"] = name
		data["content"] = content
		data["nav_id"] = navId
		data["thumb_img"] = thumbImg
		data["summary"] = summary
		data["tag"] = tag
		data["client_type"] = ClientType

		isOk := false
		if id < 1 {
			isOk = Single.AddSingle(data)
		} else {
			isOk = Single.EditSingle(id, data)
		}
		if isOk {
			return e.SUCCESS, "操作成功"
		}
		return e.ERROR, "操作失败"
	}
	return ViewErr(valid)
}
