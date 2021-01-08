package Visit

import (
	db "elearn100/Database"
	"fmt"
)

type History struct {
	Id           int    `gorm:"primary_key" json:"id"`
	Uuid         string `json:"uuid" gorm:"type:varchar(32);not null; default '';comment:'用户ID' "`
	VisitHistory string `json:"visit_history" gorm:"type:text;not null;default '';comment:'访问记录' "`
}

// @Summer 添加访问记录
func AddHistory(data map[string]interface{}) {
	result := db.Db.Create(&History{
		Uuid:         data["uuid"].(string),
		VisitHistory: data["visit_history"].(string),
	})

	if result.Error != nil {
		fmt.Printf("浏览记录失败：%s", result.Error)
	} else {
		fmt.Print("浏览记录OK")
	}
}

// @Summer 更新流量记录
func EditHistory(uuid string, updateCon map[string]interface{}) {
	result := db.Db.Model(&History{}).Where("uuid = ?", uuid).Update(updateCon)
	if result.Error != nil {
		fmt.Printf("浏览记录更新 faild：%s", result.Error)
	} else {
		fmt.Print("浏览记录更新success")
	}
}

// @Summer 获取一条记录
func GetHistory(uuid string) (his History) {
	db.Db.Where("uuid = ?", uuid).Take(&his)
	return
}
