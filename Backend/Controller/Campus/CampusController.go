package Campus

import (
	"elearn100/Model/Admin"
	"elearn100/Model/Campus"
	"elearn100/Pkg/e"
	"elearn100/Pkg/setting"
	"elearn100/Services"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

// @Summer 获取全国校区API
func GetCampus(c *gin.Context) {
	page := com.StrTo(c.Query("page")).MustInt()
	var data = make(map[string]interface{})
	data["count"] = Campus.CountCampus(data)
	data["list"] = Campus.GetCampus(page, data)
	data["size"] = setting.PageSize
	data["a_level"] = 1
	data["areas"] = Admin.GetAreas(data)
	e.Success(c, "全国校区", data)
}

// @Summer 获取单个校区API
func DetailCampus(c *gin.Context) {
	e.Success(c, "校区详情", Services.DetailCampus(c))
}

// @Summer 省统计
func GroupCampuses(c *gin.Context) {
	e.Success(c, "全国校区", Services.GroupCampus())
}

// @Summer 获取全国校区API 带缓冲区的
func GetCampuses(c *gin.Context) {
	e.Success(c, "全国校区", Services.GetCampus(c))
}

// @Summer 获取全国校区API 带缓冲区的
func AddCampuses(c *gin.Context) {
	code, msg := Services.AddCampus(c)
	e.SendRes(c, code, msg, "")
}
