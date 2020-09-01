package Elearn

import (
	"fmt"
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

// 获取elearn100 通过IP 获取浏览记录
func GetVisitByIp(ip string) (visitE JfsdVisit) {
	elearnDb.Where("ip = ?", ip).Find(&visitE)
	return
}

// elearn100 新增浏览记录
func AddVisit(data map[string]interface{}) {
	result := elearnDb.Create(&JfsdVisit{
		Uuid:       data["uuid"].(string),
		FirstUrl:   data["FirstUrl"].(string),
		Ip:         data["Ip"].(string),
		FromUrl:    data["FromUrl"].(string),
		CreateTime: time.Now().Format("2006-01-02 15:04:05"),
	})

	if result.Error != nil {
		fmt.Printf("elelarn100 浏览记录失败：%s", result.Error)
	} else {
		fmt.Print("elelarn100 浏览记录OK", data["visit_history"].(string))
	}
}

// elearn100 更新浏览记录
func UpdateVisit(ip, visitHistory string) {
	m1 := map[string]interface{}{}

	visit := GetVisitByIp(ip)
	if !strings.Contains(visit.VisitHistory, visitHistory) {
		if visit.VisitHistory == "" {
			m1["visit_history"] = visitHistory
		} else {
			m1["visit_history"] = visit.VisitHistory + "<br/>" + visitHistory
		}
	} else {
		m1["visit_history"] = visitHistory
	}

	if visitHistory != "" {
		result := elearnDb.Model(&JfsdVisit{}).Where("ip = ? ", ip).Update(m1)
		//result := elearnDb.Save(&visit)
		if result.Error != nil {
			fmt.Printf("elelarn100 浏览记录更新 faild：%s", result.Error)
		} else {
			fmt.Print("elelarn100 浏览记录更新success")
		}
	}
}
