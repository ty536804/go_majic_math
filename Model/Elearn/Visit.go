package Elearn

import (
	"elearn100/Pkg/setting"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"strings"
	"time"
)

// @Summer elearn100 浏览记录
type JfsdVisit struct {
	ID           int    `gorm:"primary_key" json:"id"`
	Uuid         string `json:"uuid" gorm:"type:uuid(32); not null; default ''; comment:'跟踪id' "`
	FirstUrl     string `json:"first_url" gorm:"type:varchar(3000); not null; default ''; comment:'访问记录' "`
	Ip           string `json:"ip" gorm:"type:varchar(100);not null; default ''; comment:'ip' "`
	FromUrl      string `json:"content" gorm:"type:from_url(2000); not null; default ''; comment:'访问来源' "`
	CreateTime   string `json:"create_time"`
	VisitHistory string `json:"visit_history" gorm:"type:text;not null; default ''" `
}

func (p *JfsdVisit) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("create_time", time.Now().Format("2006-01-02 15:04:05"))
	return nil
}

// 获取elearn100 获取浏览记录
func GetVisit(uid string) (visitE JfsdVisit) {
	elearnDb.Where("uuid = ?", uid).Find(&visitE)
	return
}

// elearn100 新增浏览记录
func AddVisit(c *gin.Context) {
	reqURI := c.Request.URL.RequestURI()
	FromUrl := setting.ReplaceSiteUrl(c.Request.Host) + reqURI //来源页
	uid, _ := c.Cookie("53gid2")
	FirstUrl := ""
	if c.Request.Referer() == "" {
		FirstUrl = c.Request.Host + reqURI //来源页
	} else {
		FirstUrl = c.Request.Referer()
	}
	result := elearnDb.Create(&JfsdVisit{
		Uuid:       uid,
		FirstUrl:   FirstUrl,
		Ip:         c.Request.RemoteAddr,
		FromUrl:    FromUrl,
		CreateTime: time.Now().Format("2006-01-02 15:04:05"),
	})

	if result.Error != nil {
		fmt.Printf("elelarn100 浏览记录失败：%s", result.Error)
	} else {
		fmt.Print("elelarn100 浏览记录OK", c.Request.Referer())
	}
}

// elearn100 更新浏览记录
func UpdateVisit(c *gin.Context) {
	m1 := map[string]interface{}{}
	uid, errUid := c.Cookie("53gid2")

	visit := GetVisit(uid)
	if !strings.Contains(visit.VisitHistory, c.Request.Referer()) {
		if visit.VisitHistory == "" {
			m1["visit_history"] = c.Request.Referer()
		} else {
			m1["visit_history"] = visit.VisitHistory + "<br/>" + c.Request.Referer()
		}
	} else {
		m1["visit_history"] = c.Request.Referer()
	}

	if errUid == nil {
		result := elearnDb.Model(&JfsdVisit{}).Where("uuid = ? ", uid).Update(m1)
		//result := elearnDb.Save(&visit)
		if result.Error != nil {
			fmt.Printf("elelarn100 浏览记录更新 faild：%s", result.Error)
		} else {
			fmt.Print("elelarn100 浏览记录更新success")
		}
	}
}
