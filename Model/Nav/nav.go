package Nav

import (
	db "elearn100/Database"
)

type Nav struct {
	db.Model

	Name string `json:"name" gorm:"type:varchar(190);not null;default '';comment:'名称'"`
	BaseUrl string `json:"base_url" gorm:"type:varchar(190);not null;default '';comment:'跳转地址'"`
	IsShow int64 `json:"is_show" gorm:"default 1;comment:'是否展示'"`
}

// @Summer 添加数据
func AddNav(data map[string]interface{}) bool  {
	info := db.Db.Create(&Nav{
		Name:data["name"].(string),
		BaseUrl:data["base_url"].(string),
		IsShow:data["is_show"].(int64),
	})

	if info.Error != nil {
		return false
	}
	return true
}

// @Summer 编辑导航
func EditNav(id int,data ...interface{}) bool {
	navInfo := db.Db.Model(Nav{}).Where("id = ?",id).Update(data)
	if navInfo.Error != nil {
		return false
	}
	return true
}

// @Summer 获取所有导航
func Navs(maps interface{}) (navs []Nav) {
	db.Db.Where(maps).Find(&navs)
	return
}

// @Summer 获取单个导航
func GetNav(id int) (navs Nav) {
	db.Db.Where("id = ?", id).Find(&navs)
	return
}