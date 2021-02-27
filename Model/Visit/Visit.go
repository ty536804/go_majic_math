package Visit

import (
	db "elearn100/Database"
	"fmt"
)

type Visit struct {
	ID         int    `gorm:"primary_key" json:"id"`
	Uuid       string `json:"uuid" gorm:"type:uuid(32); not null; default ''; comment:'跟踪id' "`
	FirstUrl   string `json:"first_url" gorm:"type:varchar(3000); not null; default ''; comment:'访问记录' "`
	Ip         string `json:"ip" gorm:"type:varchar(100);not null; default ''; comment:'ip' "`
	FromUrl    string `json:"content" gorm:"type:from_url(2000); not null; default ''; comment:'访问来源' "`
	CreateTime string `json:"create_time"`
}

// @Summer 获取单条数据
func GetVisit(uid string) (visit Visit) {
	db.Db.Select("uuid,id").Where("uuid = ?", uid).Take(&visit)
	return
}

// @Title 新增浏览记录
// @Param	data	map[string]interface
func AddVisit(visit Visit, history History) {
	if result := db.Db.Create(&visit); result.Error != nil {
		fmt.Printf("首次魔法数学 浏览记录失败：%s", result.Error)
	} else {
		AddHistory(history)
		fmt.Print("首次魔法数学 浏览记录OK")
	}
}
