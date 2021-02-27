package Services

import (
	"elearn100/Model/Admin"
	"elearn100/Model/Campus"
	"elearn100/Pkg/e"
	"encoding/json"
	"github.com/astaxie/beego/validation"
	"github.com/garyburd/redigo/redis"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

var campusKey = e.REDISKey + "campus:provinceId:"

func AddCampus(c *gin.Context) (code int, err string) {
	c.Request.Body = e.GetBody(c)

	schoolName := com.StrTo(c.PostForm("school_name")).String()
	schoolTel := com.StrTo(c.PostForm("school_tel")).String()
	workerTime := com.StrTo(c.PostForm("worker_time")).String()
	address := com.StrTo(c.PostForm("address")).String()
	schoolImg := com.StrTo(c.PostForm("school_img")).String()
	province := com.StrTo(c.PostForm("province")).MustInt()
	isShow := com.StrTo(c.PostForm("is_show")).MustInt()
	id := com.StrTo(c.PostForm("id")).MustInt()
	provinceName := com.StrTo(c.PostForm("province_name")).String()

	if code, err := validCampus(schoolName, schoolTel, workerTime, address, schoolImg, provinceName, isShow, province); code == e.ERROR {
		return code, err
	}

	campus := Campus.Campus{
		SchoolName:   schoolName,
		SchoolTel:    schoolTel,
		WorkerTime:   workerTime,
		Address:      address,
		SchoolImg:    schoolImg,
		Province:     province,
		ProvinceName: provinceName,
		IsShow:       isShow,
	}
	isOk := false
	if id < 1 {
		isOk = Campus.AddCampus(campus)
	} else {
		isOk = Campus.EditCampus(id, campus)
	}
	if isOk {
		conn := e.PoolConnect()
		defer conn.Close()

		campusKey += string(province)
		conn.Do("expire", campusKey, -1)

		return e.ReSuccess()
	}
	return e.ReError()
}

// @Desc 数据校验
func validCampus(schoolName, schoolTel, workerTime, address, schoolImg, provinceName string, isShow, province int) (int, string) {
	valid := validation.Validation{}
	valid.Required(schoolName, "school_name").Message("学校名称不能为空")
	valid.Required(schoolTel, "school_tel").Message("学校联系电话不能为空")
	valid.Required(workerTime, "worker_time").Message("学校工作日不能为空")
	valid.Required(address, "address").Message("学校地址不能为空")
	valid.Required(schoolImg, "school_img").Message("学校图片不能为空")
	valid.Required(province, "province").Message("省不能为空")
	valid.Required(provinceName, "province_name").Message("省不能为空")
	valid.Required(isShow, "is_show").Message("状态必须选择")
	if !valid.HasErrors() {
		return e.ReSuccess()
	}
	return e.ViewErr(valid)
}

// @Desc获取省份下面的校区
// @Param provinceId int 省ID
func RedisGetCampus(provinceId int) []Campus.Campus {
	conn := e.PoolConnect()
	defer conn.Close()

	campusKey += string(provinceId)
	var campuses []Campus.Campus

	if exists, _ := redis.Bool(conn.Do("exists", campusKey)); exists {
		values, _ := redis.Values(conn.Do("lrange", campusKey, 0, -1))
		var campus Campus.Campus
		for _, v := range values {
			if err := json.Unmarshal(v.([]byte), &campus); err == nil {
				campuses = append(campuses, campus)
			}
		}
	} else {
		data := make(map[string]interface{})
		data["province"] = provinceId
		data["is_show"] = 1
		campuses := Campus.GetCampus(1, data)
		for _, v := range campuses {
			if jsonStr, err := json.Marshal(v); err == nil {
				conn.Do("rpush", campusKey, jsonStr)
			}
		}
		conn.Do("expire", campusKey, e.VALIDTime)
	}
	return campuses
}

// @Summer 获取缓冲区的校区
func GetCampus(c *gin.Context) (campuses []Campus.Campus) {
	name := com.StrTo(c.PostForm("province")).String()
	res := Admin.GetArea(name)
	RedisGetCampus(res.GaodeId)
	return
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
