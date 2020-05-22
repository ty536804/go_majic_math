package Services

import (
	"elearn100/Model/Banner"
	"elearn100/Pkg/e"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"strings"
)

// @Summer 添加/编辑图片
func AddBanner(c *gin.Context) (code int, err string) {
	c.Request.Body = e.GetBody(c)

	id := com.StrTo(c.PostForm("id")).MustInt()
	bName := com.StrTo(c.PostForm("bname")).String()
	bPosition := com.StrTo(c.PostForm("bposition")).MustInt()
	imgUrl := com.StrTo(c.PostForm("imgurl")).String()
	targetLink := com.StrTo(c.PostForm("target_link")).String()
	isShow := com.StrTo(c.PostForm("is_show")).MustInt()
	info := com.StrTo(c.PostForm("info")).String()
	if strings.HasPrefix(imgUrl, "/static/upload/") {
		imgUrl = strings.Replace(imgUrl, "/static/upload/", "", -1)
	}
	valid := validation.Validation{}
	valid.Required(bName, "bname").Message("名称不能为空")
	valid.Required(bPosition, "bposition").Message("展示位置必须选择")
	valid.Required(imgUrl, "imgurl").Message("上传图片")
	valid.Required(isShow, "is_show").Message("状态必须选择")

	var data = make(map[string]interface{})
	if !valid.HasErrors() {
		data["target_link"] = targetLink
		data["info"] = info
		data["bname"] = bName
		data["bposition"] = bPosition
		data["imgurl"] = imgUrl
		data["is_show"] = isShow
		isOK := false
		if id < 1 {
			isOK = Banner.AddBanner(data)
		} else {
			isOK = Banner.EditBanner(id, data)
		}
		if isOK {
			return e.SUCCESS, "操作成功"
		}
		return e.ERROR, "操作失败"
	}
	return ViewErr(valid)
}
