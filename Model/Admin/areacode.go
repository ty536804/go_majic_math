package Admin

import (
	db "elearn100/Database"
)

//部门表
type SysAreaCode struct {
	db.Model

	Aid      int    `json:"aid" gorm:"comment:'区域编号'"`
	ALevel   int    `json:"a_level" gorm:"comment:'1:省级 2:市级 3:区级'"`
	GaodeId  int    `json:"gaode_id" gorm:"comment:'高德地图id编号'"`
	Aname    string `json:"aname" gorm:"type:varchar(100);not null;default '';comment:'区域名称'"`
	ParentId int    `json:"parent_id" gorm:"comment:'父节点'"`
	aStatus  int    `json:"a_status" gorm:"not null;default 1;comment:'1 有效 0 无效'"`
}

// @Summer获取校区列表
func GetAreas(where map[string]interface{}) (areacode []SysAreaCode) {
	db.Db.Where(where).Find(&areacode)
	return
}

// @Summer获取校区列表
func GetArea(aname string) (area SysAreaCode) {
	db.Db.Where("aname = ?", aname).First(&area)
	return
}
