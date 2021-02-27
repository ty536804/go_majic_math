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
func AddHistory(history History) {
	if result := db.Db.Create(&history); result.Error != nil {
		fmt.Printf("add 失败：%s", result.Error)
	} else {
		fmt.Print("add OK")
	}
}

// @Title 更新浏览记录
// @Param uuid string 用户ID
// @Param updateCon map[string]interface{} 更新内容
func EditHistory(uuid string, history History) {
	if result := db.Db.Model(&History{}).Where("uuid = ?", uuid).Updates(history); result.Error != nil {
		fmt.Printf("content faild：%s", result.Error)
	} else {
		fmt.Print("content success")
	}
}

// @Title  获取一条记录
// @Param  uuid string 用户ID
func GetHistory(uuid string) (his History) {
	db.Db.Where("uuid = ?", uuid).Take(&his)
	return
}
