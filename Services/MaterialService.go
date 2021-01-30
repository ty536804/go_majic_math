package Services

import (
	"elearn100/Model/Material"
	"elearn100/Pkg/e"
	"fmt"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

// @Summer 添加/编辑 视频
func AddMaterial(c *gin.Context) (code int, err string) {
	if err := c.Bind(&c.Request.Body); err != nil {
		fmt.Println("上传内容失败")
	}

	id := com.StrTo(c.PostForm("id")).MustInt()
	title := com.StrTo(c.PostForm("title")).String()
	videoSrc := com.StrTo(c.PostForm("video_src")).String()
	localSrc := com.StrTo(c.PostForm("local_src")).String()
	searCode := com.StrTo(c.PostForm("code")).String()
	isShow := com.StrTo(c.PostForm("is_show")).MustInt()
	isHot := com.StrTo(c.PostForm("is_hot")).MustInt()

	valid := validation.Validation{}
	valid.Required(title, "title").Message("标题不能为空")
	valid.Required(videoSrc, "video_src").Message("视频地址不能为空")
	valid.Required(localSrc, "local_src").Message("视频地址不能为空")
	valid.Required(isShow, "is_show").Message("请选择是否展示")
	valid.Required(isHot, "is_hot").Message("排序不能为空")

	isOk := false
	data := make(map[string]interface{})
	if !valid.HasErrors() {
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
			return e.SUCCESS, "操作成功"
		} else {
			return e.ERROR, "上传失败"
		}
	}
	return ViewErr(valid)
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
