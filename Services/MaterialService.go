package Services

import (
	"elearn100/Model/Material"
	"elearn100/Pkg/e"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

// @Summer 添加/编辑 视频
func AddMaterial(c *gin.Context) (code int, err string) {
	if err := c.Bind(&c.Request.Body); err != nil {
		return e.ReError()
	}

	id := com.StrTo(c.PostForm("id")).MustInt()
	title := com.StrTo(c.PostForm("title")).String()
	videoSrc := com.StrTo(c.PostForm("video_src")).String()
	localSrc := com.StrTo(c.PostForm("local_src")).String()
	searCode := com.StrTo(c.PostForm("code")).String()
	isShow := com.StrTo(c.PostForm("is_show")).MustInt()
	isHot := com.StrTo(c.PostForm("is_hot")).MustInt()

	data := make(map[string]interface{})
	if code, msg := validMater(title, videoSrc, localSrc, isShow, isHot); code == e.ERROR {
		return code, msg
	}
	isOk := false
	data["title"] = title
	data["video_src"] = videoSrc
	data["local_src"] = localSrc
	data["is_show"] = isShow
	data["is_hot"] = isHot
	data["code"] = searCode
	if id >= 1 {
		isOk = Material.EditMaterial(id, data)
	} else {
		isOk = Material.AddMaterial(data)
	}
	if isOk {
		return e.ReSuccess()
	}
	return e.ReError()
}

// @Desc 模型验证
func validMater(title, videoSrc, localSrc string, isShow, isHot int) (int, string) {
	valid := validation.Validation{}
	valid.Required(title, "title").Message("标题不能为空")
	valid.Required(videoSrc, "video_src").Message("视频地址不能为空")
	valid.Required(localSrc, "local_src").Message("视频地址不能为空")
	valid.Required(isShow, "is_show").Message("请选择是否展示")
	valid.Required(isHot, "is_hot").Message("排序不能为空")
	if !valid.HasErrors() {
		return e.ReSuccess()
	}
	return e.ViewErr(valid)
}

func GetMaterial(id int) Material.Material {
	return Material.GetMaterial(id)
}

func GetMaterials(page int, where map[string]interface{}) []Material.Material {
	return Material.GetMaterials(page, where)
}

func GetTotalMaterials() int {
	return Material.GetTotalMaterial()
}
