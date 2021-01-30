package Services

import (
	"elearn100/Model/Visit"
	"elearn100/Pkg/setting"
	"github.com/gin-gonic/gin"
	"strings"
)

// @Title 添加/更新 魔法数学访问记录
func AddMoFaShuXueVisit(data map[string]interface{}) {
	uid := data["uuid"].(string)
	visit := Visit.GetVisit(uid)
	if visit.ID <= 0 { //新增浏览记录
		Visit.AddVisit(data)
	} else {
		Visit.UpdateVisit(data["uuid"].(string), data["visit_history"].(string))
	}
}

// @Title 浏览记录
func AddVisit(c *gin.Context, url string) {
	reqURI := c.Request.URL.RequestURI()
	FromUrl := c.Request.Host + reqURI //来源页
	uid := strings.Split(strings.Replace(c.Request.RemoteAddr, ".", "", -1), ":")[0]
	FirstUrl := ""
	if c.Request.Referer() == "" {
		FirstUrl = setting.ReplaceSiteUrl(c.Request.Host, url, reqURI) //来源页
	} else {
		FirstUrl = c.Request.Referer()
	}
	var data = make(map[string]interface{})
	data["uuid"] = uid
	data["FirstUrl"] = FirstUrl
	data["Ip"] = strings.Split(c.Request.RemoteAddr, ":")[0]
	data["FromUrl"] = FromUrl
	data["visit_history"] = c.Request.Referer()

	if c.Request.RemoteAddr != "" {
		AddMoFaShuXueVisit(data)
	}
}
