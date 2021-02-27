package Services

import (
	"elearn100/Model/Visit"
	"elearn100/Pkg/e"
	"github.com/gin-gonic/gin"
	"strings"
	"time"
)

// @Title 添加/更新 魔法数学访问记录
func AddVisit(c *gin.Context, url string) {
	reqURI := c.Request.URL.RequestURI()
	uid := e.SubUUID(c.Request.RemoteAddr)
	visitHistory := c.Request.Referer()

	var visit Visit.Visit
	visit.Uuid = uid
	visit.FirstUrl = e.GetFirstUrl(c.Request.Referer(), c.Request.Host, url, reqURI)
	visit.FromUrl = c.Request.Host + reqURI //来源页
	visit.CreateTime = time.Now().Format("2006-01-02 15:04:05")

	var history Visit.History
	history.Uuid = uid
	history.VisitHistory = visitHistory

	if c.Request.RemoteAddr != "" {
		visit := Visit.GetVisit(uid)
		visit.Ip = strings.Split(c.Request.RemoteAddr, ":")[0]
		if visit.ID <= 0 { //新增浏览记录
			visit.Ip = strings.Split(c.Request.RemoteAddr, ":")[0]
			Visit.AddVisit(visit, history)
		} else { //更新
			visitInfo := Visit.GetHistory(uid)
			if visitInfo.VisitHistory == "" {
				visitInfo.VisitHistory = visitHistory
			} else {
				visitInfo.VisitHistory = visitInfo.VisitHistory + "<br/>" + visitHistory
			}
			Visit.EditHistory(uid, visitInfo)
		}
	}
}
