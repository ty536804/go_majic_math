package Services

import (
	"elearn100/Model/Nav"
	"elearn100/pkg/e"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

// @Summer 获取导航列表
func GetNavs(maps interface{}) (navs []Nav.Nav) {
	return Nav.Navs(maps)
}

// @Summer 获取一条导航列表
func GetNav(c *gin.Context) (navs Nav.Nav) {
	c.Request.Body = e.GetBody(c)
	id := com.StrTo(c.PostForm("id")).MustInt()
	return Nav.GetNav(id)
}

// @Summer 添加/编辑导航
func AddNav(c *gin.Context) (code int,err string) {
	c.Request.Body = e.GetBody(c)
	Name := com.StrTo(c.PostForm("name")).String()
	BaseUrl := com.StrTo(c.PostForm("base_url")).String()
	IsShow := com.StrTo(c.PostForm("is_show")).MustInt64()
	id := com.StrTo(c.PostForm("id")).MustInt()

	valid := validation.Validation{}
	valid.Required(Name,"bname").Message("名称不能为空")
	valid.Required(BaseUrl,"base_url").Message("跳转地址不能为空")
	valid.Required(IsShow,"is_show").Message("是否展示必须选择")

	var data = make(map[string]interface{})
	isOK := false
	if !valid.HasErrors() {
		data["name"] = Name
		data["base_url"] = BaseUrl
		data["is_show"] = IsShow
		if id < 1 {
			isOK = Nav.AddNav(data)
		} else {
			isOK =Nav.EditNav(id,data)
		}
		if isOK {
			return e.SUCCESS,"操作失败"
		}
		return e.ERROR,"操作失败"
	}
	return ViewErr(valid)
}