package Services

import (
	"elearn100/Model/Visit"
	"github.com/gin-gonic/gin"
)

func AddMofaShuXueVisit(c *gin.Context) {
	uid, _ := c.Cookie("53revisit")

	visit := Visit.GetVisit(uid)

	if visit.ID <= 0 { //新增浏览记录
		Visit.AddVisit(c)
	} else {
		Visit.UpdateVisit(c)
	}
}

// @Summer 浏览记录
func AddVisit(c *gin.Context) {
	AddElearnVisit(c)
	AddMofaShuXueVisit(c)
}
