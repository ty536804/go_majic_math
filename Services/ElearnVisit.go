package Services

import (
	"elearn100/Model/Elearn"
	"fmt"
	"github.com/gin-gonic/gin"
)

// @Summer elearn100 浏览历史
func AddElearnVisit(c *gin.Context) {
	uid, uOk := c.Cookie("53revisit")
	if uOk != nil {
		fmt.Print("没有缓存")
	} else {
		visit := Elearn.GetVisit(uid)
		if visit.ID > 0 {
			Elearn.UpdateVisit(c)
		} else {
			Elearn.AddVisit(c)
		}
	}
}
