package Services

import (
	"elearn100/Model/Elearn"
	"elearn100/Model/Message"
	"elearn100/Pkg/e"
	"fmt"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

// @Summer提交留言
func AddMessage(c *gin.Context) (code int, msg string) {
	var data = make(map[string]interface{})
	if err := c.Bind(&c.Request.Body); err != nil {
		fmt.Println(err)
		return e.ERROR, "操作失败"
	}

	mname := com.StrTo(c.PostForm("mname")).String()
	area := com.StrTo(c.PostForm("area")).String()
	tel := com.StrTo(c.PostForm("tel")).String()
	com := com.StrTo(c.PostForm("com")).String()

	valid := validation.Validation{}
	valid.Required(mname, "mname").Message("姓名不能为空")
	valid.Required(area, "area").Message("地区不能为空")
	valid.Required(tel, "tel").Message("选择是否展示")
	if !valid.HasErrors() {
		data["mname"] = mname
		data["area"] = area
		data["tel"] = tel
		data["content"] = ""
		data["com"] = com
		data["client"] = "pc"
		data["ip"] = c.Request.RemoteAddr
		data["channel"] = 1

		Elearn.AddMessage(c, mname, area, tel)
		if Message.AddMessage(data) {
			return e.SUCCESS, "提交成功"
		}
		return e.ERROR, "提交失败"
	}
	return ViewErr(valid)
}
