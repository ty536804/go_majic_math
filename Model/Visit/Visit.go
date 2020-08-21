package Visit

import (
	db "elearn100/Database"
	"fmt"
	"strings"
	"time"
)

type Visit struct {
	ID           int    `gorm:"primary_key" json:"id"`
	Uuid         string `json:"uuid" gorm:"type:uuid(32); not null; default ''; comment:'跟踪id' "`
	FirstUrl     string `json:"first_url" gorm:"type:varchar(3000); not null; default ''; comment:'访问记录' "`
	Ip           string `json:"ip" gorm:"type:varchar(100);not null; default ''; comment:'ip' "`
	FromUrl      string `json:"content" gorm:"type:from_url(2000); not null; default ''; comment:'访问来源' "`
	CreateTime   string `json:"create_time"`
	VisitHistory string `json:"visit_history" gorm:"type:text;not null; default ''" `
}

// @Summer 获取单条数据
func GetVisit(uid string) (visit Visit) {
	db.Db.Select("uuid,id,visit_history").Where("uuid = ?", uid).Take(&visit)
	return
}

// @Summer 新增浏览记录
func AddVisit(data map[string]interface{}) {

	result := db.Db.Create(&Visit{
		Uuid:       data["uuid"].(string),
		FirstUrl:   data["FirstUrl"].(string),
		Ip:         data["Ip"].(string),
		FromUrl:    data["FromUrl"].(string),
		CreateTime: time.Now().Format("2006-01-02 15:04:05"),
	})

	if result.Error != nil {
		fmt.Printf("brocaedu 浏览记录失败：%s", result.Error)
	} else {
		fmt.Print("brocaedu 浏览记录OK")
	}
}

// @Summer 更新数据
func UpdateVisit(uid, visitHistory string) {
	m1 := map[string]interface{}{}

	visit := GetVisit(uid)
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
		result := db.Db.Model(&Visit{}).Where("uuid = ? ", uid).Update(m1)
		if result.Error != nil {
			fmt.Printf("浏览记录更新 faild：%s", result.Error)
		} else {
			fmt.Print("浏览记录更新success")
		}
	}
}
