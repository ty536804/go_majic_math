package Services

import (
	"elearn100/Model/Admin"
	"elearn100/Model/Campus"
	"elearn100/Pkg/e"
	"encoding/json"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

func AddCampus(c *gin.Context) (code int, err string) {
	c.Request.Body = e.GetBody(c)

	schoolName := com.StrTo(c.PostForm("school_name")).String()
	schoolTel := com.StrTo(c.PostForm("school_tel")).String()
	workerTime := com.StrTo(c.PostForm("worker_time")).String()
	address := com.StrTo(c.PostForm("address")).String()
	schoolImg := com.StrTo(c.PostForm("school_img")).String()
	province := com.StrTo(c.PostForm("province")).MustInt()
	id := com.StrTo(c.PostForm("id")).MustInt()
	provinceName := com.StrTo(c.PostForm("province_name")).String()

	valid := validation.Validation{}
	valid.Required(schoolName, "school_name").Message("学校名称不能为空")
	valid.Required(schoolTel, "school_tel").Message("学校联系电话不能为空")
	valid.Required(workerTime, "worker_time").Message("学校工作日不能为空")
	valid.Required(address, "address").Message("学校地址不能为空")
	valid.Required(schoolImg, "school_img").Message("学校图片不能为空")
	valid.Required(province, "province").Message("省不能为空")
	valid.Required(provinceName, "province_name").Message("省不能为空")

	data := make(map[string]interface{})

	if !valid.HasErrors() {
		isOk := false
		data["school_name"] = schoolName
		data["school_tel"] = schoolTel
		data["worker_time"] = workerTime
		data["address"] = address
		data["school_img"] = schoolImg
		data["province"] = province
		data["province_name"] = provinceName
		if id < 1 {
			isOk = Campus.AddCampus(data)
		} else {
			isOk = Campus.EditCampus(id, data)
		}
		if isOk {
			SaveCampus(province)
			return e.SUCCESS, "操作成功"
		}
		return e.ERROR, "操作失败"
	}
	return ViewErr(valid)
}

// @Summer 缓冲区存储校区
func SaveCampus(id int) bool {
	data := make(map[string]interface{})
	data["province"] = id
	res := Campus.GetCampus(1, data)
	b, _ := json.Marshal(res)
	con := string(b)
	key := "campus" + string(id)
	isOk := e.SetMenuVal(key, con)
	return isOk
}

// @Summer 获取缓冲区的校区
func GetCampus(c *gin.Context) (campuses []Campus.Campus) {
	name := com.StrTo(c.PostForm("province")).String()
	res := Admin.GetArea(name)
	isOk, singleList := e.GetVal("campus" + string(res.GaodeId))
	if !isOk {
		isOk = SaveCampus(res.GaodeId)
		if isOk {
			json.Unmarshal([]byte(singleList), &campuses)
		}
	}
	json.Unmarshal([]byte(singleList), &campuses)
	return
}

// @Summer 获取校区
func GetCampuses(c *gin.Context) (data map[string]interface{}) {
	page := com.StrTo(c.PostForm("page")).MustInt()
	param := make(map[string]interface{})
	count := make(map[string]interface{})
	param["count"] = Campus.CountCampus(count)
	param["list"] = Campus.GetCampus(page, count)
	count["a_level"] = 1
	param["areacode"] = Admin.GetAreas(count)
	return param
}

// @Summer 获取校区
func DetailCampus(c *gin.Context) (data map[string]interface{}) {
	id := com.StrTo(c.PostForm("id")).MustInt()
	param := make(map[string]interface{})
	param["detail"] = Campus.DetailCampus(id)
	return param
}

// @Summer 省统计校区
func GroupCampus() (data map[string]interface{}) {
	param := make(map[string]interface{})
	param["detail"] = Campus.GroupCampus()
	return param
}
