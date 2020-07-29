package Services

import (
	"elearn100/Model/Elearn"
	"github.com/gin-gonic/gin"
	"strings"
)

// @Summer elearn100 浏览历史
func AddElearnVisit(c *gin.Context) {
	uid := strings.Split(strings.Replace(c.Request.RemoteAddr, ".", "", -1), ":")[0]
	visit := Elearn.GetVisit(uid)
	if visit.ID > 0 {
		Elearn.UpdateVisit(c)
	} else {
		Elearn.AddVisit(c)
	}
}
