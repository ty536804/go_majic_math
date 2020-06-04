package Services

import (
	"elearn100/Model/Single"
	"elearn100/Pkg/e"
	"encoding/json"
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

	valid := validation.Validation{}
	valid.Required(name, "name").Message("标题不能为空")
	valid.Required(content, "content").Message("内容不能为空")
	valid.Required(navId, "nav_id").Message("栏目不能为空")

	if !valid.HasErrors() {
		data["name"] = name
		data["content"] = content
		data["nav_id"] = navId
		data["thumb_img"] = thumbImg
		data["summary"] = summary

		isOk := false
		if id < 1 {
			isOk = Single.AddSingle(data)
		} else {
			isOk = Single.EditSingle(id, data)
		}
		if isOk {
			SaveSingle(navId)
			return e.SUCCESS, "操作成功"
		}
		return e.ERROR, "操作失败"
	}
	return ViewErr(valid)
}

// @Summer 存redis
func SaveSingle(id int) {
	list := Single.GetAllSingle(id)
	userResult, _ := json.Marshal(list)
	con := string(userResult)
	e.SetMenuVal("single_"+string(id), con)
	return
}

// 获取缓存中的内容
func GetSingle(id int) (single []Single.Single) {
	isOk, singleList := e.GetVal("single_" + string(id))
	if !isOk {
		SaveSingle(1)
	}
	json.Unmarshal([]byte(singleList), &single)
	return
}

// 首页接口数据
func GetCon(id int) (con map[string]interface{}) {
	tagList := Single.GetTag(id)
	data := make(map[string]interface{})
	for _, k := range tagList {
		data[k.Name] = Single.GetCon(k.Name)
	}
	return data
}
